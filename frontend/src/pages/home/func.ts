import { Message } from "@arco-design/web-vue";
export const copyList = (event: MouseEvent) => {
    // 获取dom上的data-list属性 e.g. data-list="1,2,3"  逗号后面加上空格
    const list = (event.target as HTMLElement).dataset.list?.replace(/,/g, ', ') // 不会更新
    // const list = (event.target as HTMLElement).getAttribute("list")?.replace(/,/g, ', ')
    navigator.clipboard.writeText(list || "").then(
        function () {
            /* clipboard successfully set */
            // 提示复制成功
            // alert('复制成功')
            Message.success('复制成功')
        },
        function () {
            /* clipboard write failed */
            // alert('复制失败')
            Message.error('复制失败')
        }
    );
    /**
     * 以下为弃用方法
     * document.execCommand 已被废弃
     */
    // // 创建一个input标签
    // const input = document.createElement('input')
    // // 将list赋值给input的value
    // input.value = list || ""
    // // 将input添加到body中
    // document.body.appendChild(input)
    // // 选中input的内容
    // input.select()
    // // 执行复制命令
    // document.execCommand('copy')
    // // 移除input
    // document.body.removeChild(input)
}