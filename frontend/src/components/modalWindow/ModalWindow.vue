<template>
  <div class="blackout">
    <div class="body" :style="{backgroundColor: themeStore.getColor3}">
      <div class="header" :style="{backgroundColor: themeStore.getColor4}">
        <span></span>
        <h2>{{header}}</h2>
        <cross
          @click.stop="closeModalWindow"
        />
      </div>
      <slot/>
    </div>
  </div>
</template>

<script>
import Cross from "@/components/icons/Сross.vue";
import { useThemeStore } from "@/stores/theme";

export default {
  name: "modalWindow",
  components: {
    Cross
  },
  props:{
    header:{type: String, default: ""}
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  methods:{
    closeModalWindow(){
      this.$emit("closeModalWindow")
    }
  }
}
</script>

<style scoped>
.blackout{
  position: absolute;
  z-index: 10;
  width: 100vw;
  height: 100vh;
  top: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  .body{
    width: 60%;
    height: 60%;
    .header{
      height: 80px;
      width: 100%;
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0 20px;
      box-sizing: border-box;
      span{
        width: 80px;
      }
      svg{
        height: 80px;
      }
    }
  }
}
</style>