<template>
  <div class="chart">
    <chart-item
        v-for="(date, idx) in dates"
        :key="idx"
        :value="dataset[date]"
        :previousValue="idx > 0 ? dataset[dates [idx-1]] : undefined"
        :followingValue="idx < dates.length ? dataset[dates [idx+1]] : undefined"
        :maxValue = "maxValue"
        :minValue = "minValue"
        :color="color"
        :widthColum="widthColum"
        :lowerBackgroundFill="config.lowerBackgroundFill"
        :frameColor="themeStore.getColorText"
    />
  </div>
</template>

<script>
import ChartItem from "@/components/diagram/ChartItem.vue";
import {useThemeStore} from "@/stores/theme.js";

export default {
  name: "chart",
  components: {
    chartItem: ChartItem
  },
  props:{
    widthColum:{type: Number},
    dates:{type: Array},
    dataset:{type: Object},
    maxValue:{type: Number},
    minValue:{type: Number},
    color:{type: String},
    config:{type: Object}
  },
  data() {
    return{
    }
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
}
</script>

<style scoped>
.chart{
  position: absolute;
  z-index: 3;
  left: 0;
  top: 0;
  display: flex;
  flex-direction: row-reverse;
  height: 100%;
  min-width: min-content;
  width: 100%;
}
</style>