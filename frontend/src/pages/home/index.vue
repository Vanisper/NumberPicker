<template>
  <section
    style="display: flex;flex-direction: column;flex-wrap: nowrap;justify-content: start;overflow: auto;height: 100%;">
    <div class="public">
      <!-- 总计 -->
      <div style="display: flex;flex-wrap: wrap;justify-content: center;margin: auto;">
        <span class="ball-blue" v-for="(ball, i) in allRes" :key="i">{{ ball
        }}</span>
      </div>
      <!-- 公告拖拽组件 -->
      <draggable class="list-group" :list="publicList" :group="{ name: `Category-uuid`, pull: 'clone', put: false }"
        item-key="value">
        <template #item="{ element }">
          <div class="list-group-item">
            {{ element.name }}
          </div>
        </template>
      </draggable>
    </div>
    <div style="display: flex;flex-wrap: nowrap;justify-content: start;overflow: auto;height: 100%;">
      <div class="main" v-for="(_, index) in mainBox">
        <div v-for="(_item, idx) in cards[index]" :key="idx" class="card-list">
          <card :configs="configs" :destroy="_item.destroy" :index="index" :idx="idx"
            @update:intersection="getUpdateIntersection" />
          <delete-item class="del-btn" v-if="!_item.destroy" @click="delCard(index, idx)" />
        </div>
        <div><add-item :text="'新增'" @click="cards[index].push({ list: [] })" style="margin:10px auto;" />
        </div>
        <div style="display: flex;justify-content: center;flex-wrap: wrap;max-width: 580px;">
          <span class="ball-yellow"
            v-for="(ball, i) in getIntersection(...cards[index].filter(v => !v.destroy && v.list.length).map(v => v.list))"
            :key="i">{{ ball
            }}</span>
        </div>
      </div>
    </div>
    <!-- 穿越组件  document.querySelector("#header > div > div.header-center > div")-->
    <!-- <Teleport to="#header > div > div.header-center > div">
      <span class="ball-blue-common" v-for="(ball, i) in allRes" :key="i">{{ ball
      }}</span>
    </Teleport> -->
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
<style>
.list-group-item {
  position: relative;
  display: block;
  padding: 0.75rem 1.25rem;
  margin: 0 10px;
  background-color: var(--background-primary);
  border: 1px solid var(--border-secondary);
  cursor: move;

  &:first-child {
    border-top-left-radius: inherit;
    border-top-right-radius: inherit;
  }
}
</style>
<style lang="less" scoped>
.del-btn {
  position: absolute;
  right: 10px;
  top: 10px;

  @media (max-width: 1280px) {
    // 旋转90度
    transform: rotate(-90deg);
    top: 55px;
    right: -38px;
  }
}


.public {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 5px;
  flex-direction: column;

  .list-group {
    display: flex;
  }
}


.main {
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  width: calc(calc((100% - 30px) / 3));
  border: 1px salmon solid;
  box-sizing: border-box;
  margin: 5px;
  overflow: auto;

  .card-list {
    display: flex;
    flex-direction: column;
    position: relative;
  }
}

.ball-yellow {
  --size: 30px;
  user-select: none;
  display: inline-block;
  width: var(--size);
  height: var(--size);
  line-height: var(--size);
  font-size: 15px;
  font-weight: bolder;
  border-radius: 50%;
  // background: #ff755b;
  // 渐变
  background-image: radial-gradient(circle at 0% 0%, #ffef77 0, #ffde66 16.67%, #ffc950 33.33%, #f6b135 50%, #e99917 66.67%, #e08300 83.33%, #da7100 100%);
  // 阴影
  box-shadow: 0px 5px 10px 0px rgb(0 0 0 / 28%);
  color: #fff;
  margin: 10px;
  text-align: center;
}

.ball-blue {
  --size: 30px;
  user-select: none;
  display: inline-block;
  width: var(--size);
  height: var(--size);
  line-height: var(--size);
  font-size: 15px;
  font-weight: bolder;
  border-radius: 50%;
  // background: #ff755b;
  // 渐变
  background-image: linear-gradient(180deg, #9792f4 0, #7680ec 25%, #486ee5 50%, #005ede 75%, #004fd6 100%);
  // 阴影
  box-shadow: 0px 5px 10px 0px rgb(0 0 0 / 28%);
  color: #fff;
  margin: 10px;
  text-align: center;
}

.ball-blue-common {
  --size: 23px;
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
  background-image: linear-gradient(180deg, #9792f4 0, #7680ec 25%, #486ee5 50%, #005ede 75%, #004fd6 100%);
  color: #fff;
  margin: 10px;
  text-align: center;
}
</style>
