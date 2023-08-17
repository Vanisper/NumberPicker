<template>
  <section>
    <!--ABC共同号码区域-->
    <!-- v-if="allRes.length>0" -->
    <div class="bg-gray-600/20 rounded-3xl m-5 py-5 px-5  flex justify-start flex-wrap items-center gap-2 bg-gray-200 ">
      <div>
      </div>
      <div class="flex-1 flex justify-start items-center gap-3 flex-wrap">
        <span class="ball-blue" v-for="(ball, i) in allRes" :key="i">
          {{ ball }}
        </span>
        <!--<span class="ball-blue" v-for="(ball, i) in 49" :key="i">-->
        <!--  {{ ball }}-->
        <!--</span>-->
        <span class="text-center block w-full text-2xl text-gray-600" v-if="allRes.length == 0">
          暂无全局相同的号码
        </span>
      </div>
      <div class="text-center p-5">
        <!--<div class="mb-2 text-gray-700">-->
        <!--  <p>全局横向</p>-->
        <!--  <p>相同号码</p>-->
        <!--</div>-->
        <div :data-list="allRes" class="daya-copy-btn cursor-pointer" @click="copyList">
          复制
        </div>
      </div>
    </div>

    <!--分类选择区域-->
    <div class="bg-gray-600/20 rounded-3xl m-5 py-5 px-5 ">
      <!--分类拖拽器-->
      <draggable class="flex gap-2 flex-wrap justify-between items-center" :list="publicList"
        :group="{ name: `Category-uuid`, pull: 'clone', put: false }" item-key="value">
        <template #item="{ element }">
          <div class="daya-class-btn cursor-move">
            {{ element.name }}
          </div>
        </template>
      </draggable>
    </div>

    <!--复制3个一模一样的拖动盒子-->
    <div class=" m-5 flex justify-between items-start gap-10">

      <div class="bg-blue-100 border-blue-200 border rounded-3xl p-5 flex-1 flex flex-col" v-for="(_, index) in mainBox">
        <!--<div class="flex justify-between items-center">-->
        <!--  <div>{{ index+1 }}</div>-->
        <!--  <div class="daya-copy-btn">-->
        <!--    复制-->
        <!--  </div>-->
        <!--</div>-->

        <div v-for="(_item, idx) in cards[index]" :key="idx" class="bg-blue-200 rounded-2xl mb-2">
          <card :configs="configs" :destroy="_item.destroy" :index="index" :idx="idx"
            @update:intersection="getUpdateIntersection" />
          <div>
            <delete-item class="del-btn" v-if="!_item.destroy" @click="delCard(index, idx)" />
          </div>
        </div>

        <div class="mt-3">
          <add-item :text="'新增'" @click="cards[index].push({ list: [] })" />
        </div>

        <div class="bg-yellow-100 rounded-2xl mt-5 py-2 ">
          <div class="flex justify-between items-center w-full px-5">
            <div class="flex-1 flex justify-start items-center flex-wrap gap-2">
              <span class="ball-yellow"
                v-for="(ball, i) in getIntersection(...cards[index].filter(v => !v.destroy && v.list.length).map(v => v.list))"
                :key="i">
                {{ ball }}
              </span>

              <div class="text-gray-500"
                v-if="getIntersection(...cards[index].filter(v => !v.destroy && v.list.length).map(v => v.list)).length == 0">
                本列竖向没有相同号码
              </div>
            </div>
            <div :data-list="getIntersection(...cards[index].filter(v => !v.destroy && v.list.length).map(v => v.list))" class="daya-copy-btn ml-5 cursor-pointer" @click="copyList">
              复制
            </div>
          </div>
        </div>

      </div>
    </div>

  </section>
</template>

// https://github.com/vuejs/core/issues/7799
// 解决了setup组件中没有显式定义组件name的问题  路由缓存的include、exclude属性匹配的是组件的name属性
<script lang="ts">
import { PageEnum } from "@/enums/pageEnum";

export default {
  name: PageEnum.BASE_HOME_NAME,
};
</script>
<script lang="ts" setup>
import { ExportConfig } from "@wailsjs/go/main/App";
import { EventsOn } from "@wailsjs/runtime";
import { computed, onMounted, ref } from "vue";
import card from "./includes/card.vue";
import { IConfig } from "@/types/config.interface";
import AddItem from "@/components/button/AddItem.vue";
import DeleteItem from "@/components/button/DeleteItem.vue";
import { copyList } from "./func";

import draggable from "vuedraggable";

const configs = ref<IConfig>();

const getConfig = async () => {
  await ExportConfig().then((res: IConfig) => {
    configs.value = res;
  }).catch((err) => {
    console.log(err);
  })
}
const publicList = computed(() => Object.keys(configs.value?.Category || {}).map((key) => {
  return {
    name: key,
    value: configs.value?.Category[key],
  }
}))

onMounted(async () => {
  await getConfig();
  EventsOn("update:config", async (data) => {
    configs.value = data;
  })
})

// 创建3个主盒子
interface ICard {
  order?: number;
  destroy?: boolean;
  list: number[]
}

const mainBox = ref(Array.from({ length: 3 }))
const cards = ref<ICard[][]>(mainBox.value.map(() => {
  return [{
    list: []
  }]
}))

// 多个一维数组的交集
function getIntersection<T>(...arrays: T[][]): T[] {
  if (arrays.length === 0) {
    return [];
  }

  const result: T[] = [];

  for (const item of arrays[0]) {
    if (result.includes(item)) {
      continue;
    }

    let isCommon = true;

    for (let i = 1; i < arrays.length; i++) {
      if (!arrays[i].includes(item)) {
        isCommon = false;
        break;
      }
    }

    if (isCommon) {
      result.push(item);
    }
  }

  return result;
}

const delCard = (index: number, idx: number) => {
  // 弹窗 是否确认删除
  // window.confirm("确认删除？") && (cards.value[index][idx].destroy = true)
  cards.value[index][idx].destroy = true;
}
// @ts-ignore
const isAllDestroy = computed(() => {
  return cards.value.map(value => value.map(v => v.destroy ? true : false)).flatMap(v => v).every(v => v)
})
const getUpdateIntersection = (intersection: number[], index: number, idx: number) => {
  cards.value[index][idx].list = intersection
}
const allRes = computed(() => {
  const res = getIntersection(...cards.value.filter(v => v.length && !v.map(v => v.destroy).every(v => v)).map(v => getIntersection(...v.filter(v => !v.destroy && v.list.length).map(v => v.list))))

  return res;
})
</script>

<style lang="less">
.daya-class-btn {
  @apply inline-block text-center bg-blue-400 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded;
}

.daya-copy-btn {
  @apply h-8 max-h-8 w-20 text-xs text-center bg-gray-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded;
}
</style>

<style scoped lang="less">
.ball-yellow {
  @apply text-white font-bold rounded-full w-8 h-8 flex items-center justify-center m-1;
  background-image: radial-gradient(circle at 0% 0%, #ffef77 0, #ffde66 16.67%, #ffc950 33.33%, #f6b135 50%, #e99917 66.67%, #e08300 83.33%, #da7100 100%);
  box-shadow: 0px 5px 10px 0px rgb(0 0 0 / 28%);
}

.ball-blue {
  @apply text-white font-bold rounded-full w-8 h-8 flex items-center justify-center m-1;
  background-image: linear-gradient(180deg, #9792f4 0, #7680ec 25%, #486ee5 50%, #005ede 75%, #004fd6 100%);
  box-shadow: 0px 5px 10px 0px rgb(0 0 0 / 28%);
}
</style>
