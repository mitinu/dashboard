<template>
  <header :style="{backgroundColor: themeStore.getColor4}">
    <div class="header">
      <div class="container">
        <div>
          <menu-burger-horizontal-icon
              @click.stop="switchVisibilityBurgerMenu"
              :color="themeStore.getColorText"
          />
        </div>
        <div class="title"><h3>{{title}}</h3></div>
      </div>
      <div class="container">
<!--        <div class="text">-->
<!--          <div><span>источники информации</span></div>-->
<!--          <arrow-down-->
<!--              :color="themeStore.getColorText"-->
<!--          />-->
<!--        </div>-->
        <div><button-exit/></div>
      </div>
    </div>
  </header>
</template>

<script>
import arrowDown from "@/components/icons/ArrowDown.vue";
import menuBurgerHorizontalIcon from "@/components/icons/MenuBurgerHorizontalIcon.vue";
import buttonExit from "@/components/UI/ButtonExit.vue";
import { useThemeStore } from "@/stores/theme";

export default {
  name: "mainHeader",
  components: {
    menuBurgerHorizontalIcon,
    arrowDown,
    buttonExit
  },
  props:{
    title:{type: String, default: "title"}
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  methods:{
    switchVisibilityBurgerMenu(){
      this.$emit("switchVisibilityBurgerMenu")
    }
  }
}
</script>

<style scoped>
  header{
    align-content: end;
    padding: 20px;
    height: 50px;
    .header{
      display: flex;
      align-items: center;
      justify-content: space-between;
      .title{
        margin-left: 80px;
      }
    }
    .container{
      padding-left: 50px;
      display: grid;
      grid-auto-flow: column;
      gap: 30px;
      align-content: center;
      .text{
        display: flex;
        align-items: center;
        svg{
          width: 20px;
          height: 20px;
        }
      }
      svg{
        width: 35px;
        height: 35px;
      }
    }
  }
</style>