<template>
  <div class="buttonReturn">
    <left-arrow class="icon"/><span> Назад</span>
  </div>
</template>

<script>
import LeftArrow from "@/components/icons/LeftArrow.vue";
export default {
  name: "buttonReturn",
  components: {LeftArrow}
}
</script>

<style scoped>
.buttonReturn{
  background-color: #42425c;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
  padding: 10px 20px;
  border-radius: 5%;
}
.icon{
  width: 20px;
  height: 20px;
}
</style>