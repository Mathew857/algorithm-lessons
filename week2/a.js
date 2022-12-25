foo(1);
function foo(i) {
    if (i == 4) {
        return;
    }
    console.log('foo() begin:' + i);
    foo(i + 1);
    console.log('foo() end:' + i);
}
// 输出 ? 

