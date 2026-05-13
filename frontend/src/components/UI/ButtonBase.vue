<template>
  <input
      type="button"
      class="buttonBase"
      :value=value
  >
</template>

<script>
export default{
  name:"buttonBase",
  props:{
    value:{
      type:String,
    },
  }
}
</script>

<style>
.buttonBase{
  background-color: #2f2f40;
  height: 40px;
  padding: 0 40px;
  border-radius: 20px ;
  color: #FFFFFF;
  cursor: pointer;
}
</style>