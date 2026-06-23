<template>
  <div class="header">
    <div class="title"><h3>{{title}}</h3></div>
    <legend-chart
        v-if="legend.length>0"
        :symbols="legend"
    />

  </div>

</template>

<script>
import legendChart from "@/components/diagram/LegendChart.vue";

export default {
  name: "headerDiagramPage",
  components: {
    legendChart
  },
  props:{
    legend:{type:Array},
    title:{type: String}
  }
}
</script>
<style scoped>
.header{
  display: flex;
  gap: 50px;
  margin-bottom: 40px;
  height: 50px;
}
</style>