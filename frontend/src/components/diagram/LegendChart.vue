<template>
  <div class="legendChart">
    <div
      class="symbol"
      v-for="(symbol, idx) in symbols"
      :key="idx"
    >
      <div class="title"><span>{{symbol.title}}</span></div>
      <div class="color" :style="{ backgroundColor: symbol.color }"></div>
    </div>
  </div>
</template>

<script>
export default {
  props:{
    symbols:{type: Array}
  },
  name: "legendChart",
}
</script>

<style scoped>
.legendChart{
  display: flex;
  flex-wrap: wrap;
  gap: 30px;
  .symbol{
    display: flex;
    gap: 10px;
    .color{
      width: 23px;
      height: 23px;
      border-radius: 100%;
    }
  }
}
</style>