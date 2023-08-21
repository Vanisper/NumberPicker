<template>
    <div class="card" style=" height: 100%; " v-if="!destroy && configs?.Category">
        <div class="left " style=" height: 100%;">
            <div style="display: flex; height: 100%; justify-content: space-around; ">
                <div class="drag-box daya-same-class-box">
                    <!--左侧分类框-->
                    <!-- 预留6个高度的位置  或者去掉具体数字  直接预留所有可能的高度 -->
                    <draggable class="list-group"
                        :style="{ 'min-height': (6 || Object.keys(configs.Category).length) * 43 + 'px' }" :list="list1"
                        :group="`Category-${uid}`" @change="onAdd1" @start="canRemove = true" @end="canRemove = false"
                        item-key="value">
                        <template #item="{ element }">
                            <div class="daya-class-btn mb-0.5">
                                {{ element.name }}
                            </div>
                        </template>
                    </draggable>
                </div>
                <div class="drag-box daya-same-class-box public" style="display: none;">
                    <!--分类选择框 隐藏-->
                    <draggable class="list-group" :list="publicList"
                        :group="{ name: `Category-${uid}`, pull: 'clone', put: false }" item-key="value">
                        <template #item="{ element }">
                            <div class="daya-class-btn mb-0.5">
                                {{ element.name }}
                            </div>
                        </template>
                    </draggable>
                </div>
                <div class="drag-box daya-same-class-box">
                    <!--右侧分类框-->
                    <draggable class="list-group" :list="list2" :group="`Category-${uid}`" @change="onAdd2"
                        @start="canRemove = true" @end="canRemove = false" item-key="value">
                        <template #item="{ element }">
                            <div class="daya-class-btn mb-0.5">
                                <pre>{{ element.name }}</pre>
                            </div>
                        </template>
                    </draggable>
                </div>
            </div>

            <div v-if="canRemove" class="trash-box" :style="{ backgroundColor: canRemove ? '#ff000069' : '' }">
                <draggable class="list-group" :list="trashList" :group="`Category-${uid}`" item-key="value"
                    @change="remove">
                    <template #item></template>
                </draggable>
                <div
                    style="position: absolute;top: 50%;left: 50%;transform: translate(-50%,-50%);color: #fff;font-size: 15px;font-weight: bolder;width: 100%;display: flex;justify-content: center;">
                    拖向此处，移除之
                </div>
            </div>
        </div>
        <div class="right" style="flex:1; height: 100%;position: relative;align-items: center;">
            <span v-for="(item, index) in listIntersection" :key="index" class="ball" :style="{
                'background-image': callback.map(v => v.value).includes(item.value) ? 'radial-gradient(circle at 0% 0%, #ffef77 0, #ffde66 16.67%, #ffc950 33.33%, #f6b135 50%, #e99917 66.67%, #e08300 83.33%, #da7100 100%)' : '',
                display: (!showPublicBall && callback.map(v => v.value).includes(item.value)) ? 'none' : ''
            }">{{ item.value }}</span>
            <div v-if="listIntersection.length == 0" style="color: #666;  height: 100%;">无相同号码</div>
            <div :data-list="fixListIntersection.filter(v => !v.public).map(v => v.value)" v-else
                class="daya-copy-btn cursor-pointer ml-5" @click="copyList">
                复制
            </div>
            <div style="display: flex;">
                <div class="daya-copy-btn cursor-pointer ml-5" style="background-color: rgb(64, 219, 109);" @click="paste">
                    粘贴
                </div>
                <div class="daya-copy-btn cursor-pointer ml-5" style="background-color: rgb(243, 89, 94);" @click="clean">
                    清空
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { IConfig } from "@/types/config.interface";
import { computed, ref, watch } from "vue";
import { Message } from "@arco-design/web-vue";
import { copyList } from "../func"
import draggable from "vuedraggable";
import { v4 as uuidv4 } from 'uuid';
// 确保不会跨组件拖拽 （停用  转为公共分类组件）
const uid = "uuid" || uuidv4();

const props = defineProps({
    configs: {
        type: Object as () => IConfig,
        default: undefined,
    },
    index: {
        type: Number,
        required: true
    },
    idx: {
        type: Number,
        required: true
    },
    destroy: {
        type: Boolean,
        default: false,
    },
    callback: {
        type: Array as () => {
            value: number;
            from: number[];
        }[],
        required: true
    },
    showPublicBall: {
        type: Boolean
    }
})
const enums = defineEmits(['update:intersection', 'remove:category'])
interface IList {
    name: string;
    value: number[];
}
const configs = ref<IConfig>();
const publicList = computed(() => Object.keys(configs.value?.Category || {}).map((key) => {
    return {
        name: key,
        value: configs.value?.Category[key],
    }
}))
const canRemove = ref<boolean>(false);
//@ts-ignore 不绑定list就可以充当垃圾桶的功能
const trashList = ref<IList[]>([]);
const list1 = ref<IList[]>([]);
const list2 = ref<IList[]>([]);
// list的值的交集
const listIntersection = computed(() => {
    // const uniqueArr = Array.from(new Set(arr.flatMap((item) => item)));

    const list1Value = list1.value.map((item) => item.value);
    const list2Value = list2.value.map((item) => item.value);
    const list1ValueFlat = Array.from(new Set(list1Value.flatMap((item) => item)));
    const list2ValueFlat = Array.from(new Set(list2Value.flatMap((item) => item)));
    const intersection = list1ValueFlat.filter((item) => list2ValueFlat.includes(item)).sort((a, b) => a - b).map(v => ({
        value: v,
        public: false,
    }));
    const res = Array.from(new Set([...intersection, ...customList.value].map(v => v.value))).sort((a, b) => a - b).map(v => ({
        value: v,
        public: false,
    }))

    enums('update:intersection', res, props.index, props.idx)
    return res;
})

const fixListIntersection = computed(() => {
    const list = listIntersection.value;
    props.callback.map(v => v.value).forEach(v => {
        const index = (list.map(v => v.value)).indexOf(v)
        if (index != -1) {
            list[index].public = true
        }
    })

    return list
})

const onAdd1 = (_evt: any) => {
    if (list1.value?.length && list1.value.length > 1) {
        const allName = list1.value.map((item) => item.name);
        if (!_evt.added) {
            return;
        }
        const currentElement = _evt.added.element;
        const currentName = currentElement.name;
        const currentIndex = _evt.added.newIndex;
        // console.log(currentName,currentIndex);
        // 如果当前元素的名称在数组中有多个
        if (allName.filter((item) => item === currentName).length > 1) {
            list1.value.splice(currentIndex, 1);
        }
    }
}

const onAdd2 = (_evt: any) => {
    if (list2.value?.length && list2.value.length > 1) {
        const allName = list2.value.map((item) => item.name);
        if (!_evt.added) {
            return;
        }
        const currentElement = _evt.added.element;
        const currentName = currentElement.name;
        const currentIndex = _evt.added.newIndex;
        // console.log(currentName,currentIndex);
        // 如果当前元素的名称在数组中有多个
        if (allName.filter((item) => item === currentName).length > 1) {
            list2.value.splice(currentIndex, 1);
        }
    }
}

const remove = (_evt: any) => {
    enums("remove:category", _evt.added.element, props.index, props.idx)
    trashList.value = []
}
const customList = ref<{
    public: boolean,
    value: number
}[]>([])
const paste = () => {

    if (navigator.clipboard) {
        navigator.clipboard.readText()
            .then(text => {
                const content = Array.from(new Set(text.match(/\d+/g)));
                if (content.length == 0) {
                    return Message.error('剪切板并没有检测到数字')
                }
                Message.success('粘贴 ' + content.join(","))
                const list = content.map(v => ({
                    value: +v,
                    public: false
                }));
                list && (customList.value = list);
            })
            .catch(error => console.log('获取剪贴板内容失败：', error));
    } else {
        console.log('当前浏览器不支持Clipboard API');
    }
}
const clean = () => {
    list1.value = [];
    list2.value = [];
    customList.value = [];
    Message.success("已清空")
}

watch(() => props.configs, (value) => {
    configs.value = value;
    // 更新一下list1、list2的值
    if (value?.Category) {
        list1.value = list1.value.map((item) => {
            return {
                name: item.name,
                value: value.Category[item.name],
            }
        })
        list2.value = list2.value.map((item) => {
            return {
                name: item.name,
                value: value.Category[item.name],
            }
        })
    }
}, {
    immediate: true,
})

</script>

<style lang="less" scoped>
.card {
    display: flex;
    flex-direction: row;
    align-items: center;

    @media (max-width: 1280px) {
        flex-direction: column;
    }
}



.left {
    display: flex;
    flex-direction: column;
    position: relative;
    align-items: center;
}



.ball {

    @apply text-white font-bold rounded-full w-8 h-8 flex items-center justify-center m-1;
    background-image: radial-gradient(circle at 12.5% 12.5%, #ffb476 0, #ffa470 16.67%, #ff8f67 33.33%, #ff745b 50%, #f05851 66.67%, #e53e4e 83.33%, #dd2550 100%);
    box-shadow: 0px 5px 10px 0px rgb(0 0 0 / 28%);
}

.daya-same-class-box {
    min-height: 100%;
}


.right {
    border: 1px solid var(--border-primary);
    padding: 5px;
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    justify-content: center;
    flex-grow: 1;
}



.drag-box {
    width: 100px;
    border: 1px solid var(--border-primary);
    margin: 10px;
    padding: 10px;
    overflow: auto;
}


.public {
    background: #eee;
}

.list-group {
    height: 100%;
}

.list-group-item {
    position: relative;
    display: block;
    padding: 10px;
    background-color: var(--background-primary);
    border: 1px solid var(--border-secondary);
    cursor: move;
    display: flex;
    justify-content: center;
    align-items: center;
}



.trash-box {
    width: calc(100% - 20px);
    height: 100px;
    position: absolute;
    bottom: 10px;
}
</style>
