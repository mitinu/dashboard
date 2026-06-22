<template>
  <div class="comparisonHistogram">
    <comparison-histogram-item
        v-for="(date, idx) in dates"
        :key="idx"
        :values="dataset[date]"
        :maxValue = "maxValue"
        :minValue = "minValue"
        :colors="colors"
        :widthColum="widthColum"
    />
  </div>
</template>

<script>
import comparisonHistogramItem from "./ComparisonHistogramItem.vue";

export default {
  name: "comparisonHistogram",
  components: {
    comparisonHistogramItem
  },
  props:{
    widthColum:{type: Number},
    dates:{type: Array},
    dataset:{type: Object},
    maxValue:{type: Number},
    minValue:{type: Number},
    ratioFillingWidth:{type: Array},
    colors:{type: Array}
  },
  data() {
    return{
    }
  },
}
</script>

<style scoped>
.comparisonHistogram{
  width: 100%;
  position: absolute;
  z-index: 2;
  left: 0;
  top: 0;
  display: flex;
  flex-direction: row-reverse;
  height: 100%;
  min-width: min-content;
}
</style>