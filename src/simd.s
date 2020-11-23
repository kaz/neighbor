	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 11, 0	sdk_version 11, 0
	.intel_syntax noprefix
	.globl	_lookup_popcnt          ## -- Begin function lookup_popcnt
	.p2align	4, 0x90
_lookup_popcnt:                         ## @lookup_popcnt
## %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	dword ptr [r8], 0
	test	esi, esi
	je	LBB0_5
## %bb.1:
	movsxd	r9, esi
	movsxd	r10, ecx
	xor	esi, esi
	xor	eax, eax
	jmp	LBB0_2
	.p2align	4, 0x90
LBB0_4:                                 ##   in Loop: Header=BB0_2 Depth=1
	inc	rax
	cmp	rax, r9
	jae	LBB0_5
LBB0_2:                                 ## =>This Inner Loop Header: Depth=1
	mov	rcx, qword ptr [rdi + 8*rax]
	xor	rcx, rdx
	popcnt	rcx, rcx
	cmp	rcx, r10
	jg	LBB0_4
## %bb.3:                               ##   in Loop: Header=BB0_2 Depth=1
	inc	esi
	mov	dword ptr [r8], esi
	jmp	LBB0_4
LBB0_5:
	mov	rsp, rbp
	pop	rbp
	ret
                                        ## -- End function
	.section	__TEXT,__const
	.p2align	5               ## -- Begin function lookup_simd
LCPI1_0:
	.space	32,15
LCPI1_1:
	.byte	0                       ## 0x0
	.byte	1                       ## 0x1
	.byte	1                       ## 0x1
	.byte	2                       ## 0x2
	.byte	1                       ## 0x1
	.byte	2                       ## 0x2
	.byte	2                       ## 0x2
	.byte	3                       ## 0x3
	.byte	1                       ## 0x1
	.byte	2                       ## 0x2
	.byte	2                       ## 0x2
	.byte	3                       ## 0x3
	.byte	2                       ## 0x2
	.byte	3                       ## 0x3
	.byte	3                       ## 0x3
	.byte	4                       ## 0x4
	.byte	0                       ## 0x0
	.byte	1                       ## 0x1
	.byte	1                       ## 0x1
	.byte	2                       ## 0x2
	.byte	1                       ## 0x1
	.byte	2                       ## 0x2
	.byte	2                       ## 0x2
	.byte	3                       ## 0x3
	.byte	1                       ## 0x1
	.byte	2                       ## 0x2
	.byte	2                       ## 0x2
	.byte	3                       ## 0x3
	.byte	2                       ## 0x2
	.byte	3                       ## 0x3
	.byte	3                       ## 0x3
	.byte	4                       ## 0x4
	.section	__TEXT,__text,regular,pure_instructions
	.globl	_lookup_simd
	.p2align	4, 0x90
_lookup_simd:                           ## @lookup_simd
## %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	dword ptr [r8], 0
	test	esi, esi
	je	LBB1_4
## %bb.1:
	vmovq	xmm0, rdx
	vpbroadcastq	ymm0, xmm0
	movsxd	rax, esi
	inc	ecx
	movsxd	rcx, ecx
	vmovq	xmm1, rcx
	vpbroadcastq	ymm1, xmm1
	xor	ecx, ecx
	vmovdqa	ymm2, ymmword ptr [rip + LCPI1_0] ## ymm2 = [15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15]
	vmovdqa	ymm3, ymmword ptr [rip + LCPI1_1] ## ymm3 = [0,1,1,2,1,2,2,3,1,2,2,3,2,3,3,4,0,1,1,2,1,2,2,3,1,2,2,3,2,3,3,4]
	vpxor	xmm4, xmm4, xmm4
	xor	edx, edx
	.p2align	4, 0x90
LBB1_2:                                 ## =>This Inner Loop Header: Depth=1
	vpxor	ymm5, ymm0, ymmword ptr [rdi + 8*rdx]
	vpsrlw	ymm6, ymm5, 4
	vpand	ymm5, ymm5, ymm2
	vpshufb	ymm5, ymm3, ymm5
	vpand	ymm6, ymm6, ymm2
	vpshufb	ymm6, ymm3, ymm6
	vpaddb	ymm5, ymm6, ymm5
	vpsadbw	ymm5, ymm5, ymm4
	vpcmpgtb	ymm5, ymm1, ymm5
	vpmovmskb	esi, ymm5
	popcnt	rsi, rsi
	add	ecx, esi
	add	rdx, 4
	cmp	rdx, rax
	jb	LBB1_2
## %bb.3:
	mov	dword ptr [r8], ecx
LBB1_4:
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
                                        ## -- End function
.subsections_via_symbols
