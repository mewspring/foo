declare void @a()
declare void @b()

define void @f() {
  call void @a()
  call void @b()
  ret void
}
