<template>
<div class="BurgerMenu">
  <menu-burger-horizontal-icon
      @click.stop="switchVisibilityBurgerMenu"
  />
  <router-link to="/operationalInformation">оперативные сведения</router-link>
  <router-link to="/monthlyInformation">Ежемесячная информация</router-link>
</div>
</template>

<script>
import menuBurgerHorizontalIcon from "@/components/icons/menuBurgerHorizontalIcon.vue";

export default {
  name: "BurgerMenu",
  components: {menuBurgerHorizontalIcon},
  methods:{
    switchVisibilityBurgerMenu(){
      this.$emit("switchVisibilityBurgerMenu")
    }
  }
}
</script>

<style scoped>
.BurgerMenu{
  position: fixed;
  z-index: 100;
  height: 100%;
  top: 0;
  left: 0;
  width: 110px;
  background-color: #2f2f40;
  padding: 30px 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  svg{
    width: 100%;
    height: 35px;
    margin-bottom: 50px;
  }
}
</style>