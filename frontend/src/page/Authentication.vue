<template>
  <div class="authentication">
    <div class="container">
      <h1>Вход</h1>
      <div class="inputs">
        <input-text-base
            placeholder="Логин"
            v-model=login
        />
        <span v-if="errors.login" class="error-text">{{ errors.login }}</span>
        <input-text-base
            placeholder="Пароль"
            v-model=password
        />
        <span v-if="errors.password" class="error-text">{{ errors.password }}</span>
        <a @click.stop="this.$emit('confirmationAuthentication')">Забыл(а) пароль</a>
      </div>
      <button-base value="Войти" @click.stop="authorization" />
    </div>
  </div>
</template>

<script>
import inputTextBase from "@/components/UI/InputTextBase.vue"
import buttonBase from "@/components/UI/ButtonBase.vue";
export default {
  name: "authentication",
  components:{
    inputTextBase,
    buttonBase
  },
  data() {
    return {
      login: "",
      password: "",
      errors: {
        login: null,
        password: null
      }
    }
  },
  methods:{
    validate() {
      this.errors = { login: null, password: null };
      let isValid = true;

      if (!this.login) {
        this.errors.login = "Логин обязателен для заполнения";
        isValid = false;
      } else if (this.login.length < 3) {
        this.errors.login = "Логин должен быть не менее 3-х символов";
        isValid = false;
      }

      if (!this.password) {
        this.errors.password = "Введите пароль";
        isValid = false;
      } else if (this.password.length < 6) {
        this.errors.password = "Пароль слишком короткий (мин. 6 символов)";
        isValid = false;
      }

      return isValid;
    },
    authorization(){
      if (this.validate()) {
        this.$emit("confirmationAuthentication")
      }
    }
  }
}
</script>

<style scoped>
.error-text {
  color: #ff4d4d;
  font-size: 14px;
  margin-top: -15px; /* Чтобы текст был ближе к инпуту */
  align-self: flex-start;
}
.authentication{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  .container{
    min-width: 450px;
    min-height: 350px;
    width: 70%;
    height: 60%;
    background-color: #171620;


    border: 2px solid #2f2f40;
    border-radius: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 50px;

    display: flex;
    .inputs{
      width: 375px;
      display: flex;
      flex-direction: column;
      gap: 20px;
      a{
        cursor: pointer;
        user-select: none;
        min-width: 40%;
        width: min-content;
      }
    }
    input[type=button]{
      font-size: 24px;
    }
  }
}
</style>