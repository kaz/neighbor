find.s: src/find_simd.s find.go
	docker run --rm -v $(PWD):/work c2goasm -a $< $@

src/find_simd.s: src/find_simd.c
	clang -S -Ofast -march=native -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -o $@ $<
