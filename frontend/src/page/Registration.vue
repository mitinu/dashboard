<template>
  <div class="registration">
    <div class="container" :style="{backgroundColor: themeStore.getColor3}">
      <h1>Регистрация</h1>
      <div class="inputs">
        <div class="field">
          <input-text-base
              placeholder="Логин"
              v-model="form.login"
          />
          <span v-if="errors.login" class="error-message">{{ errors.login }}</span>
        </div>

        <div class="field">
          <SelectBase
              placeholder="Уровень доступа"
              :options="form.accessStatus"
              v-model="form.levelAccess"
          />
        </div>

        <div class="field">
          <input-text-base
              placeholder="Пароль"
              type="password"
              v-model="form.password"
          />
          <span v-if="errors.password" class="error-message">{{ errors.password }}</span>
        </div>

        <div class="field">
          <input-text-base
              placeholder="Повторите пароль"
              type="password"
              v-model="form.passwordRepeat"
          />
          <span v-if="errors.passwordRepeat" class="error-message">{{ errors.passwordRepeat }}</span>
        </div>
      </div>

      <button-base value="Создать аккаунт" @click.stop="register" />
    </div>
  </div>
</template>

<script>
import inputTextBase from "@/components/UI/InputTextBase.vue"
import buttonBase from "@/components/UI/ButtonBase.vue";
import SelectBase from "@/components/UI/selectBase.vue";
import { useThemeStore } from "@/stores/theme";

export default {
  name: "Registration",
  components: {
    SelectBase,
    inputTextBase,
    buttonBase
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  data() {
    return {
      titleHeader: "Регистрация новых пользователей",
      form: {
        login: "",
        password: "",
        passwordRepeat: "",
        accessStatus: [{value:2, title:"админ"},{value:1, title:"пользователь"}],
        levelAccess: null
      },
      errors: {}
    }
  },
  methods: {
    validate() {
      this.errors = {};

      if (!this.form.login) {
        this.errors.login = "Придумайте логин";
      } else if (this.form.login.length < 3) {
        this.errors.login = "Минимум 3 символа";
      }

      if (!this.form.password) {
        this.errors.password = "Введите пароль";
      } else if (this.form.password.length < 6) {
        this.errors.password = "Пароль должен быть от 6 символов";
      }

      if (this.form.password !== this.form.passwordRepeat) {
        this.errors.passwordRepeat = "Пароли не совпадают";
      }

      return Object.keys(this.errors).length === 0;
    },

    register() {
      if (this.validate()) {
        this.$emit("submitRegistration");
      }
    }
  },
  mounted() {
    this.$emit("setTitleMain", "")
    this.$emit("setTitleHeader", this.titleHeader)
    this.$emit("setVisibilityButtonReturnMain", false)
  },
}
</script>

<style scoped lang="scss">
.registration {
  display: flex;
  justify-content: center;
  align-items: center;
  height: calc(100% - 110px);

  .container {
    min-width: 450px;
    padding: 40px;
    border: 2px solid #2f2f40;
    border-radius: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 30px;

    .inputs {
      width: 100%;
      max-width: 375px;
      display: flex;
      flex-direction: column;
      gap: 15px;

      .field {
        display: flex;
        flex-direction: column;
        gap: 5px;

        .error-message {
          color: #ff4d4d;
          font-size: 12px;
          padding-left: 5px;
        }
      }
    }

    .back-link {
      cursor: pointer;
      color: #71717a;
      font-size: 14px;
      text-decoration: underline;
      &:hover {
        color: #fff;
      }
    }
  }
}
</style>