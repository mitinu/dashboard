<template>
  <header :style="{backgroundColor: themeStore.getColor4}">
    <div class="header">
      <burger-menu
          :class="{ 'is-open': visibilityBurgerMenu }"
          @switchVisibilityBurgerMenu="switchVisibilityBurgerMenu"
      />
      <div class="container">
        <div>
          <menu-burger-horizontal-icon
              @click.stop="switchVisibilityBurgerMenu"
              class="positionAbsolute"
              :color="themeStore.getColorText"
          />
        </div>
        <div class="title"><h3>{{title}}</h3></div>
      </div>
      <div class="container">
        <div class="text">
          <div><span>источники информации</span></div>
          <arrow-down
              :color="themeStore.getColorText"
          />
        </div>
        <div><button-exit/></div>
      </div>
    </div>
  </header>
</template>

<script>
import arrowDown from "@/components/icons/ArrowDown.vue";
import menuBurgerHorizontalIcon from "@/components/icons/MenuBurgerHorizontalIcon.vue";
import buttonExit from "@/components/UI/ButtonExit.vue";
import burgerMenu from "@/components/BurgerMenu.vue";
import { useThemeStore } from "@/stores/theme";

export default {
  name: "mainHeader",
  components: {
    menuBurgerHorizontalIcon,
    arrowDown,
    buttonExit,
    burgerMenu
  },
  props:{
    title:{type: String, default: "title"}
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  data(){
    return{
      visibilityBurgerMenu: false
    }
  },
  methods:{
    switchVisibilityBurgerMenu(){
      this.visibilityBurgerMenu = !this.visibilityBurgerMenu
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
  :deep(.BurgerMenu) {
    transform: translateX(-100%);
    transition: transform 0.3s ease-in-out;
  }
  :deep(.BurgerMenu.is-open) {
    transform: translateX(0);
  }
</style>