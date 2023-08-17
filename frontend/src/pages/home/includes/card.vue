<template>
    <div class="card" v-if="!destroy && configs?.Category">
        <div class="left">
            <div style="display: flex;">
                <div class="drag-box">
                    <!-- 预留6个高度的位置  或者去掉具体数字  直接预留所有可能的高度 -->
                    <draggable class="list-group"
                        :style="{ 'min-height': (6 || Object.keys(configs.Category).length) * 43 + 'px' }" :list="list1"
                        :group="`Category-${uid}`" @change="onAdd1" @start="canRemove = true" @end="canRemove = false"
                        item-key="value">
                        <template #item="{ element }">
                            <div class="list-group-item">
                                {{ element.name }}
                            </div>
                        </template>
                    </draggable>
                </div>
                <div class="drag-box public" style="display: none;">
                    <draggable class="list-group" :list="publicList"
                        :group="{ name: `Category-${uid}`, pull: 'clone', put: false }" item-key="value">
                        <template #item="{ element }">
                            <div class="list-group-item">
                                {{ element.name }}
                            </div>
                        </template>
                    </draggable>
                </div>
                <div class="drag-box">
                    <draggable class="list-group" :list="list2" :group="`Category-${uid}`" @change="onAdd2"
                        @start="canRemove = true" @end="canRemove = false" item-key="value">
                        <template #item="{ element }">
                            <div class="list-group-item">
                                <pre>{{ element.name }}</pre>
                            </div>
                        </template>
                    </draggable>
                </div>
            </div>
            <div v-if="canRemove" class="trash-box" :style="{ backgroundColor: canRemove ? '#ff000069' : '' }">
                <draggable class="list-group" :group="`Category-${uid}`" item-key="value">
                    <template #item></template>
                </draggable>
                <div
                    style="position: absolute;top: 50%;left: 50%;transform: translate(-50%,-50%);color: #fff;font-size: 15px;font-weight: bolder;width: 100%;display: flex;justify-content: center;">
                    拖向此处，移除之
                </div>
            </div>
        </div>
        <div class="right">
            <span v-for="(item, index) in listIntersection" :key="index" class="ball">{{ item }}</span>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { IConfig } from "@/types/config.interface";
import { computed, ref, watch } from "vue";
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
})
const enums = defineEmits(['update:intersection'])
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
    const intersection = list1ValueFlat.filter((item) => list2ValueFlat.includes(item)).sort((a, b) => a - b);
    enums('update:intersection', intersection, props.index, props.idx)
    return intersection;
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
// watch(() => props.destroy, (value) => {
//     if (value) {
//         list1.value = [];
//         list2.value = [];
//     }
// }, {
//     immediate: true,
// })
</script>

<style lang="less" scoped>
.card {
    display: flex;
    flex-direction: row;
    align-items: center;

    @media (max-width: 1280px) {
        flex-direction: column;
    }

    .left {
        display: flex;
        flex-direction: column;
        position: relative;
        align-items: center;

        .trash-box {
            width: calc(100% - 20px);
            height: 100px;
            position: absolute;
            bottom: 10px;
            // transform: translateY(100%);


            .list-group {
                height: 100%;
                overflow: auto;

                .list-group-item {
                    position: relative;
                    display: block;
                    padding: 0.75rem 1.25rem;
                    background-color: var(--background-primary);
                    border: 1px solid var(--border-secondary);
                    cursor: move;

                    &:first-child {
                        border-top-left-radius: inherit;
                        border-top-right-radius: inherit;
                    }
                }
            }
        }


        .drag-box {
            width: 100px;
            border: 1px solid var(--border-primary);
            margin: 10px;
            padding: 10px;
            overflow: auto;

            &.public {
                background: #eee;
            }

            .list-group {
                height: 100%;

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

                    &:first-child {
                        border-top-left-radius: inherit;
                        border-top-right-radius: inherit;
                    }
                }
            }

        }
    }

    .right {
        // max-width: 220px;
        border: 1px solid var(--border-primary);
        padding: 5px;
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        justify-content: center;
        flex-grow: 1;

        .ball {
            --size: 25px;
            user-select: none;
            display: inline-block;
            width: var(--size);
            height: var(--size);
            line-height: var(--size);
            font-size: 12px;
            font-weight: bolder;
            border-radius: 50%;
            // background: #ff755b;
            // 渐变
            background-image: radial-gradient(circle at 12.5% 12.5%, #ffb476 0, #ffa470 16.67%, #ff8f67 33.33%, #ff745b 50%, #f05851 66.67%, #e53e4e 83.33%, #dd2550 100%);
            // 阴影
            box-shadow: 0px 5px 10px 0px rgb(0 0 0 / 28%);
            color: #fff;
            margin: 10px;
            text-align: center;
        }
    }
}
</style>
