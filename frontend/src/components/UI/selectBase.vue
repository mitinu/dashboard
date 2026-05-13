<template>
  <select
      class="selectBase"
      :value="modelValue"
      @change="$emit('update:modelValue', $event.target.value)"
  >
    <option
        v-for="option in options"
        :key="option.value"
        :value="option.value"
    >
      {{ option.title }}
    </option>
  </select>
</template>

<script>
export default {
  name: "SelectBase",
  props: {
    modelValue: {type: [String, Number], default: ""},
    options: {type: Array, default: () => []}
  },
  emits: ['update:modelValue']
}
</script>

<style>
.selectBase {
  background-color: #2f2f40;
  Width: 395px;
  Height: 50px;
  border-radius: 20px;
  color: #FFFFFFCC;
  font-size: 20px;
  font-weight: 500;
  padding-left: 20px;
}


.selectBase option {
  background-color: #2f2f40;
  color: #FFFFFF;
}
</style>
