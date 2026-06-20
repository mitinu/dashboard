<template>
  <div class="BurgerMenu" :class="{ 'is-open': isOpen }" :style="menuStyles">
    <div class="content">
      <div class="body">
        <router-link class="menu-link" to="/registration">
          <span class="link-text">Регистрация</span>
        </router-link>
        <router-link class="menu-link" to="/users">
          <span class="link-text">Пользователи</span>
        </router-link>
        <router-link class="menu-link" to="/operationalInformation">
          <span class="link-text">Оперативные сведения</span>
        </router-link>
        <router-link class="menu-link" to="/monthlyInformation">
          <span class="link-text">Экономические данные</span>
        </router-link>
      </div>

      <div class="footer">
        <gears
            class="iconGears"
            @click.stop="openModalWindow"
            :color="themeStore.getColorText"
        />
      </div>

    </div>
    <modal-window
        v-if="visibilityModalWindow"
        @closeModalWindow="closeModalWindow"
        :header="headerModalWindow"
    >
      <setting @setHeader="setHeaderModalWindow" />
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
  mixins: [
    mixinModalWindow
  ],
  props: {
    isOpen: {
      type: Boolean,
      default: false
    }
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  computed: {
    menuStyles() {
      return {
        '--menu-bg': this.themeStore.getColor4,
        '--menu-text': this.themeStore.getColorText,
      };
    }
  }
}
// TODO добавить скрол в случи когда список страничек будет велик
</script>

<style scoped>
.BurgerMenu {
  width: 0;
  transition: width 0.3s ease-in-out;
  overflow: hidden;

  /* Переменные по умолчанию на случай отсутствия стора */
  --menu-bg: #1e1e1e;
  --menu-text: #ffffff;
}

.BurgerMenu.is-open {
  width: 175px;
}

.content {
  width: 175px;
  height: 100%;
  background-color: var(--menu-bg);
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  transform: translateX(-100%);
  transition: transform 0.3s ease-in-out;
}

.BurgerMenu.is-open .content {
  transform: translateX(0);
}

.body {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 10px;
  align-items: center;
}

.menu-link {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: var(--menu-text);
  font-size: 10px;
  font-weight: 500;
  border-radius: 8px;
  transition: background-color 0.2s ease, transform 0.1s ease;
  width: 80%;
  padding: 5px 10px;
  box-sizing: border-box;
}

.menu-link:hover {
  background-color: color-mix(in srgb, var(--menu-text) 10%, transparent);
}

.menu-link.router-link-active {
  background-color: color-mix(in srgb, var(--menu-text) 15%, transparent);
  font-weight: 600;
}

.footer {
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
  svg{
    height: 40px;
    width: 40px;
  }
}

.iconGears {
  height: 28px;
  width: 28px;
  cursor: pointer;
  opacity: 0.8;
  transition: transform 0.3s ease, opacity 0.2s ease;
}

.iconGears:hover {
  opacity: 1;
  transform: rotate(45deg);
}
</style>