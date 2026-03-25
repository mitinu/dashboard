<template>
  <div class="informationBlock">
    <div><h3>{{title}}</h3></div>
    <div><span class="finePrint">данные актуальны на {{dateRange.start}} - {{dateRange.end}}</span></div>
  </div>
</template>

<script>
export default {
  name: "headerInformationBlock",
  props:{
    title:{type: String},
    dateRange:{type: Object}
  }
}
</script>

<style scoped>
.informationBlock{
  padding: 10px 20px;
  border-bottom: 1px solid #4b4b4b;
}
</style>