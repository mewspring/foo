@s = constant [4 x i8] c"foo\00"

define void @f() {
	%1 = getelementptr [4 x i8], [4 x i8]* @s, i64 0, i64 0
	ret void
}
