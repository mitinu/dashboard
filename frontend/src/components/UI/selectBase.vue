<template>
  <select
      class="selectBase"
      :class="{ 'has-placeholder': !modelValue && placeholder }"
      :value="modelValue"
      @change="$emit('update:modelValue', $event.target.value)"
      :style="{
        backgroundColor: themeStore.getColor4,
        color: themeStore.getColorText
      }"
  >
    <option v-if="placeholder" value="" disabled selected hidden>{{ placeholder }}</option>
    <option
        v-for="option in options"
        :key="option.value"
        :value="option.value"
        :style="{
          backgroundColor: themeStore.getColor4,
          color: themeStore.getColorText
        }"
    >
      {{ option.title }}
    </option>
  </select>
</template>

<script>
import { useThemeStore } from "@/stores/theme"

export default {
  name: "SelectBase",
  props: {
    modelValue: {type: [String, Number], default: ""},
    options: {type: Array, default: () => []},
    placeholder: {type: String, default: ""}
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  emits: ['update:modelValue']
}
</script>

<style>
.selectBase {
  Width: 395px;
  Height: 50px;
  border-radius: 20px;
  font-size: 20px;
  font-weight: 500;
  padding-left: 20px;
}

.selectBase.has-placeholder {
  color: #757575; /* Полупрозрачный белый (или любой твой цвет для подсказок) */
}

/* Возвращает нормальный цвет для остальных опций внутри списка */
.selectBase option {
  color: #FFFFFF;
}
</style>
