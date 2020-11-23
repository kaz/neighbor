#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <immintrin.h>

volatile int data_len = 1000000;
volatile uint64_t* data;

volatile int bench_data_len = 100;
volatile int bench_tolerance = 5;

///////////////////////////////////////////////////////////////////////////////////////////////////

uint64_t popcnt(uint64_t n) {
	n = (n & 0x5555555555555555) + ((n >>  1) & 0x5555555555555555);
	n = (n & 0x3333333333333333) + ((n >>  2) & 0x3333333333333333);
	n = (n & 0x0f0f0f0f0f0f0f0f) + ((n >>  4) & 0x0f0f0f0f0f0f0f0f);
	n = (n & 0x00ff00ff00ff00ff) + ((n >>  8) & 0x00ff00ff00ff00ff);
	n = (n & 0x0000ffff0000ffff) + ((n >> 16) & 0x0000ffff0000ffff);
	n = (n & 0x00000000ffffffff) + ((n >> 32) & 0x00000000ffffffff);
	return n;
}

int find_v10(uint64_t value, int tolerance) {
	int result = 0;
	for (int i = 0; i < data_len; i++) {
		if (popcnt(value ^ data[i]) <= tolerance) {
			result++;
		}
	}
	return result;
}

///////////////////////////////////////////////////////////////////////////////////////////////////

int find_v11(uint64_t value, int tolerance) {
	int result = 0;
	for (int i = 0; i < data_len; i++) {
		if (_mm_popcnt_u64(value ^ data[i]) <= tolerance) {
			result++;
		}
	}
	return result;
}

///////////////////////////////////////////////////////////////////////////////////////////////////

int find_v12(uint64_t value, int tolerance) {
	int result = 0;

	__m256i zero = _mm256_setzero_si256();
	__m256i mask = _mm256_set1_epi8(0xF);

	__m256i population = _mm256_setr_epi8(
		0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
		0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4
	);

	__m256i values = _mm256_set1_epi64x(value);
	__m256i tolerances = _mm256_set1_epi64x(tolerance+1);

	for (int i = 0; i < data_len; i += 4) {
		__m256i vec = _mm256_xor_si256(_mm256_setr_epi64x(data[i], data[i+1], data[i+2], data[i+3]), values);
		__m256i vec_lower = _mm256_and_si256(vec, mask);
		__m256i vec_upper = _mm256_and_si256(_mm256_srli_epi16(vec, 4), mask);

		vec_lower = _mm256_shuffle_epi8(population, vec_lower);
		vec_upper = _mm256_shuffle_epi8(population, vec_upper);
		vec = _mm256_sad_epu8(_mm256_add_epi8(vec_lower, vec_upper), zero);

		int flag = _mm256_movemask_epi8(_mm256_cmpgt_epi8(tolerances, vec));
		result += _mm_popcnt_u64(flag);
	}

	return result;
}

///////////////////////////////////////////////////////////////////////////////////////////////////

typedef struct node {
	uint64_t value;
	struct node** children;
} node_t;

node_t* root = NULL;

node_t* new_node(uint64_t value) {
	node_t* node = malloc(sizeof(node_t));
	node->value = value;
	node->children = calloc(64, sizeof(node_t*));
	return node;
}

void add_node(node_t* node, uint64_t value) {
	int d = _mm_popcnt_u64(node->value ^ value);
	if (node->children[d] == NULL) {
		node->children[d] = new_node(value);
	} else {
		add_node(node->children[d], value);
	}
}

void build_v20() {
	root = new_node(data[0]);
	for (int i = 1; i < data_len; i++) {
		add_node(root, data[i]);
	}
}

int _find_v20(node_t* node, uint64_t value, int tolerance) {
	int d = _mm_popcnt_u64(node->value ^ value);
	int result = d <= tolerance;

	for (int i = d - tolerance; 0 <= i && i < 64 && i <= d + tolerance; i++) {
		if (node->children[i] != NULL) {
			result += _find_v20(node->children[i], value, tolerance);
		}
	}
	return result;
}
int find_v20(uint64_t value, int tolerance) {
	return _find_v20(root, value, tolerance);
}

///////////////////////////////////////////////////////////////////////////////////////////////////

uint64_t random64() {
	return (random() << 62) + (random() << 31) + random();
}

int main() {
	srandom(time(NULL));

	data = malloc(data_len * sizeof(uint64_t));
	for (int i = 0; i < data_len; i++) {
		data[i] = random64();
	}

	uint64_t* bench_data = malloc(bench_data_len * sizeof(uint64_t));
	for (int i = 0; i < bench_data_len; i++) {
		bench_data[i] = random64();
	}

	printf("---- PREPARE ----\n");
	clock_t c;

	c = clock();
	build_v20();
	printf("v20: %.2lfms\n", (double)(clock() - c) / CLOCKS_PER_SEC * 1000);

	uint64_t test = random64();
	printf("---- SANITY CHECK ----\n");
	printf("v20: %d\n", find_v20(test, 16));
	printf("v12: %d\n", find_v12(test, 16));
	printf("v11: %d\n", find_v11(test, 16));
	printf("v10: %d\n", find_v10(test, 16));

	printf("---- BENCHMARK ----\n");
	volatile int _;

	c = clock();
	for (int i = 0; i < bench_data_len; i++) {
		_ = find_v20(bench_data[i], bench_tolerance);
	}
	printf("v20: %.2lfms\n", (double)(clock() - c) / CLOCKS_PER_SEC * 1000);

	c = clock();
	for (int i = 0; i < bench_data_len; i++) {
		_ = find_v12(bench_data[i], bench_tolerance);
	}
	printf("v12: %.2lfms\n", (double)(clock() - c) / CLOCKS_PER_SEC * 1000);

	c = clock();
	for (int i = 0; i < bench_data_len; i++) {
		_ = find_v11(bench_data[i], bench_tolerance);
	}
	printf("v11: %.2lfms\n", (double)(clock() - c) / CLOCKS_PER_SEC * 1000);

	c = clock();
	for (int i = 0; i < bench_data_len; i++) {
		_ = find_v10(bench_data[i], bench_tolerance);
	}
	printf("v10: %.2lfms\n", (double)(clock() - c) / CLOCKS_PER_SEC * 1000);
}
