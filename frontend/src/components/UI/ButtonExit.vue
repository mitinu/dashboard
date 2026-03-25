<template>
  <div>
    <button class="buttonExit"><exit-right/></button>
  </div>
</template>

<script>
import exitRight from "@/components/icons/ExitRight.vue";
export default {
  name: "buttonExit",
  components: {
    exitRight
  }
}
</script>

<style scoped>
.buttonExit{
  background-color: #42425c;
  padding: 15px;
  border-radius: 100%;
  svg{
    width: 25px;
    height: 25px;
  }
}
</style>