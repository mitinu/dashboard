<template>
  <input
      type="button"
      class="buttonBase"
      :value=value
      :style="{
        backgroundColor: themeStore.getColor4,
        color: themeStore.getColorText
      }"
  >
</template>

<script>
import { useThemeStore } from "@/stores/theme";

export default{
  name:"buttonBase",
  props:{
    value:{
      type:String,
    },
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
}
</script>

<style>
.buttonBase{
  height: 40px;
  padding: 0 40px;
  border-radius: 20px ;
  cursor: pointer;
}
</style>