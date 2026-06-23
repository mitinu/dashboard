<template>
  <div>
    <button
        class="buttonExit"
        :style="{backgroundColor: themeStore.getColor4}"
    >
      <exit-right
          :color="themeStore.getColorText"
      />
    </button>
  </div>
</template>

<script>
import exitRight from "@/components/icons/ExitRight.vue";
import {useThemeStore} from "@/stores/theme.js";

export default {
  name: "buttonExit",
  components: {
    exitRight
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
}
</script>

<style scoped>
.buttonExit{
  background-color: #42425c;
  padding: 15px;
  border-radius: 100%;
  cursor: pointer;
  svg{
    width: 25px;
    height: 25px;
  }
}
</style>