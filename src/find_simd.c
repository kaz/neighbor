#include <immintrin.h>

void find(uint64_t* haystack, int haystack_len, uint64_t needle, int tolerance, int* result) {
	__m256i zero = _mm256_setzero_si256();
	__m256i mask = _mm256_set1_epi8(0xF);

	__m256i population = _mm256_setr_epi8(
		0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
		0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4
	);

	__m256i needles = _mm256_set1_epi64x(needle);
	__m256i tolerances = _mm256_set1_epi64x(tolerance+1);

	*result = 0;
	for (uint64_t i = 0; i < haystack_len; i += 4) {
		__m256i vec = _mm256_xor_si256(_mm256_setr_epi64x(haystack[i], haystack[i+1], haystack[i+2], haystack[i+3]), needles);
		__m256i vec_lower = _mm256_and_si256(vec, mask);
		__m256i vec_upper = _mm256_and_si256(_mm256_srli_epi16(vec, 4), mask);

		vec_lower = _mm256_shuffle_epi8(population, vec_lower);
		vec_upper = _mm256_shuffle_epi8(population, vec_upper);
		vec = _mm256_sad_epu8(_mm256_add_epi8(vec_lower, vec_upper), zero);

		uint32_t flag = _mm256_movemask_epi8(_mm256_cmpgt_epi8(tolerances, vec));
		*result += _mm_popcnt_u64(flag);
	}
}

// #include <stdio.h>

// int main() {
// 	uint64_t r;
// 	uint64_t data[] = {1, 2, 3, 4, 5, 6, 7, 8};
// 	find(data, 8, 4, 1, &r);
// 	printf("%llu\n", r);
// }
