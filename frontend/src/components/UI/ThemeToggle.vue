<template>
  <label class="theme-switch" aria-label="Переключить тему">
    <input
        type="checkbox"
        v-model="isDark"
    />
    <span class="slider">
      <span class="icon-wrapper">
        <sun v-if="!isDark" class="icon sun" />
        <moon v-else class="icon moon" />
      </span>
    </span>
  </label>
</template>

<script>
import sun from "@/components/icons/Sun.vue";
import moon from "@/components/icons/Moon.vue";

export default {
  name: "themeToggle",
  components: {
    moon,
    sun,
  },
  data(){
    return{
      isDark: false
    }
  },
}
</script>

<style scoped>
.theme-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.theme-switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 32px;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0; left: 0; right: 0; bottom: 0;
  background-color: #ccc;
  transition: background-color 0.3s ease;
  border-radius: 34px;
}

.icon-wrapper {
  position: absolute;
  height: 24px;
  width: 24px;
  left: 4px;
  bottom: 4px;
  border-radius: 50%;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  display: flex;
  align-items: center;
  justify-content: center;

  box-shadow: 0 2px 4px rgba(0,0,0,0.2);
}

.icon {
  width: 16px;
  height: 16px;
  transition: transform 0.3s ease, color 0.3s ease;
}


input:checked + .slider {
  background-color: #2c3e50;
}

input:checked + .slider .icon-wrapper {
  transform: translateX(28px);
}
</style>