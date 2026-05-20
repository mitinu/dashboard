<template>
  <div class="setting">
      <div class="cell">
        <h2>тема</h2>
        <h3>светлая/темная</h3>
        <theme-toggle @change="themeToggle"/>
      </div>
  </div>
</template>

<script>
import themeToggle from "@/components/UI/ThemeToggle.vue";
import { useThemeStore } from "@/stores/theme";

export default {
  name: "setting",
  components: {
    themeToggle
  },
  data(){
    return{
      header:" настройки",
    }
  },
  methods:{
    themeToggle(event) {
      const store = useThemeStore();
      if (event.target.checked) {
        store.setBackgroundMain("#0b0b10");
        store.setColorText("#F0F8FFFF");
        store.setColor3("#171620");
        store.setColor4("#2f2f40");
      } else {
        store.setBackgroundMain("#F0F8FFFF");
        store.setColorText("#0b0b10");
        store.setColor3("#E8E6F5");
        store.setColor4("#D0D0E0");
      }
    }
  },
  mounted() {
    this.$emit("setHeader", this.header)
  }
}
// TODO добавить изменение размера h2, h3, span
</script>

<style scoped>
.setting{
  width: 100%;
  height: 100%;
  display: grid;
  .cell{
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 10px;
  }
}
</style>