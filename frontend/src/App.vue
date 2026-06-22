<template>
  <div
    class="app"
    :style="{
      backgroundColor: themeStore.getBackgroundMain,
      color: themeStore.getColorText
    }"
  >
    <div v-if="authentication">
      <main-header
          :title="titleHeader"
          @switchVisibilityBurgerMenu="switchVisibilityBurgerMenu"
      />
      <div class="app__content">
        <burger-menu
            :class="{ 'is-open': visibilityBurgerMenu }"
            :is-open="visibilityBurgerMenu"
            @switchVisibilityBurgerMenu="switchVisibilityBurgerMenu"
        />
        <main-body
            @setTitleHeader="setTitleHeader"
        />
      </div>
    </div>
    <div v-else>
      <authentication
          @confirmationAuthentication="confirmationAuthentication"
      />
    </div>
  </div>
</template>

<script>
import mainHeader from "@/components/MainHeader.vue";
import mainBody from "@/components/MainBody.vue";
import Authentication from "@/page/Authentication.vue";
import { useThemeStore } from "@/stores/theme";
import BurgerMenu from "@/components/BurgerMenu.vue";

export default {
  name:"app",
  components:{
    burgerMenu: BurgerMenu,
    Authentication,
    mainHeader,
    mainBody
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  data(){
    return{
      titleHeader:"",
      authentication: false,
      visibilityBurgerMenu: false
    }
  },
  methods:{
    setTitleHeader(title){
      this.titleHeader = title
    },
    confirmationAuthentication(){
      this.authentication = true
    },
    switchVisibilityBurgerMenu(){
      this.visibilityBurgerMenu = !this.visibilityBurgerMenu
    }
  }
}
//TODO переделать графики переделать стилизацию
</script>

<style>

.gradientLine {
  height: 1px;
  background: linear-gradient(90deg, rgba(96, 96, 96, 0.3), rgba(96, 96, 96, 0.6), rgba(96, 96, 96, 0.3));
}
h1{
  font-size: 26pt;
}
h2{
  font-size: 22pt;
}
h3{
  font-size: 18pt;
}
h4{
  font-size: 16pt;
}
span{
  white-space: pre-line;
}
.mt20{
  margin-top: 20px;
}
.mt25{
  margin-top: 20px;
}
*{
  border: none;
  margin: 0;
  padding: 0;
  font-size: 13pt;
  font-family: Calibri;
}
a{
  text-decoration: none;
}
input::placeholder {
  color: #757575;
  opacity: 1;
}
.app__content{
  display: grid;
  grid-template-columns: max-content 1fr;
  height: 100%;
}
</style>
