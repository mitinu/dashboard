<template>
  <div class="histogram">
    <histogram-item
      v-for="(date, idx) in dates"
      :key="idx"
      :value="dataset[date]"
      :maxValue = "maxValue"
      :minValue = "minValue"
      :color="color"
      :widthColum="widthColum"
    />
  </div>
</template>

<script>
import HistogramItem from "./HistogramItem.vue";

export default {
  name: "histogram",
  components: {HistogramItem},
  props:{
    widthColum:{type: Number},
    dates:{type: Array},
    dataset:{type: Object},
    maxValue:{type: Number},
    minValue:{type: Number},
    ratioFillingWidth:{type: Array},
    color:{type: String}
  },
  data() {
    return{
    }
  },
}
</script>

<style scoped>
.histogram{
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