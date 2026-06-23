<template>
  <input
      :type="type"
      class="inputTextBase"
      :placeholder=placeholder
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      :list="list"
      @change="emitChange"
      :style="{
        backgroundColor: themeStore.getColor4,
        color: themeStore.getColorText
      }"
  >
</template>

<script>
import { useThemeStore } from "@/stores/theme";

export default{
  name:"inputTextBase",
  props:{
    placeholder:{type:String},
    modelValue:{type: String},
    list:{type: String},
    type:{type: String, default: "text"}
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  methods:{
    emitChange(){
      this.$emit("change")
    }
  }
}
</script>

<style>
.inputTextBase{
  Width: 375px;
  Height: 50px;
  border-radius: 20px;
  color: #FFFFFFCC;
  font-size: 20px;
  font-weight: 500;
  padding-left: 20px;
}
</style>