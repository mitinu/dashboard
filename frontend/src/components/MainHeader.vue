<template>
  <header>
    <div class="header">
      <burger-menu
          @switchVisibilityBurgerMenu="switchVisibilityBurgerMenu"
          v-if="visibilityBurgerMenu"
      />
      <div class="container">
        <div>
          <menu-burger-horizontal-icon
              @click.stop="switchVisibilityBurgerMenu"
              class="positionAbsolute"
          />
        </div>
        <div class="title"><h3>{{title}}</h3></div>
      </div>
      <div class="container">
        <div class="text">
          <div><span>источники информации</span></div>
          <arrow-down/>
        </div>
        <div><button-exit/></div>
      </div>
    </div>
  </header>
</template>

<script>
import arrowDown from "@/components/icons/ArrowDown.vue";
import menuBurgerHorizontalIcon from "@/components/icons/menuBurgerHorizontalIcon.vue";
import buttonExit from "@/components/UI/ButtonExit.vue";
import burgerMenu from "@/components/BurgerMenu.vue";

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
    background-color: #2f2f40;
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
</style>