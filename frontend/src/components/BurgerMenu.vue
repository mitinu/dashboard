<template>
<div class="BurgerMenu" :style="{backgroundColor: themeStore.getColor4}">
  <menu-burger-horizontal-icon
      class="iconBurgerMenu"
      @click.stop="switchVisibilityBurgerMenu"
      :color="themeStore.getColorText"
  />
  <div class="body">
    <router-link :style="{color: themeStore.getColorText}" to="/operationalInformation">оперативные сведения</router-link>
    <router-link :style="{color: themeStore.getColorText}" to="/monthlyInformation">Ежемесячная информация</router-link>
    <router-link :style="{color: themeStore.getColorText}" to="/registration">регистрация</router-link>

  </div>
  <gears
      class="iconGears"
      @click.stop="openModalWindow"
      :color="themeStore.getColorText"
  />
  <modal-window
    v-if="visibilityModalWindow"
    @closeModalWindow="closeModalWindow"
    :header="headerModalWindow"
  >
    <setting
        @setHeader="setHeaderModalWindow"
    />
  </modal-window>
</div>
</template>

<script>
import menuBurgerHorizontalIcon from "@/components/icons/MenuBurgerHorizontalIcon.vue";
import Gears from "@/components/icons/Gears.vue";
import modalWindow from "@/components/modalWindow/ModalWindow.vue";
import setting from "@/page/Setting/Setting.vue";
import mixinModalWindow from "@/components/modalWindow/mixinModalWindow.vue";
import { useThemeStore } from "@/stores/theme";

export default {
  name: "BurgerMenu",
  components: {
    setting,
    modalWindow,
    Gears,
    menuBurgerHorizontalIcon
  },
  mixins:[
      mixinModalWindow
  ],
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  methods:{
    switchVisibilityBurgerMenu(){
      this.$emit("switchVisibilityBurgerMenu")
    },
  }
}
// TODO добавить скрол в случи когда список страничек будет велик
</script>

<style scoped>
.BurgerMenu{
  position: fixed;
  z-index: 100;
  top: 0;
  left: 0;
  width: 150px;
  height: 100%;
  padding: 30px 20px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  align-items: center;
  .iconBurgerMenu{
    height: 35px;
  }
  .body{
    width: 120px;
    display: flex;
    flex-direction: column;
    height: calc(100% - 70px);
    gap: 10px;
    padding: 30px 0;
    box-sizing: border-box;
  }
  .iconGears{
    height: 35px;
  }
}
</style>