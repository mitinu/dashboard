<template>
  <input
      type="text"
      class="inputTextBase"
      :placeholder=placeholder
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      :list="list"
      @change="emitChange"
  >
</template>

<script>
export default{
  name:"inputTextBase",
  props:{
    placeholder:{
      type:String
    },
    modelValue:{
      type: String
    },
    list:{
      type: String
    }
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
  background-color: #2f2f40;
  Width: 375px;
  Height: 50px;
  border-radius: 20px;
  color: #FFFFFFCC;
  font-size: 20px;
  font-weight: 500;
  padding-left: 20px;
}
</style>