simd.s: src/simd.s simd.go
	docker run --rm -v $(PWD):/work c2goasm -a $< $@

src/simd.s: src/simd.c
	clang -S -Ofast -march=native -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -o $@ $<
