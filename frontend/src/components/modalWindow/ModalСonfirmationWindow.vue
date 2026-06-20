<template>
  <div class="modal-overlay">
    <div class="modal-content" :style="{backgroundColor: themeStore.getColor4}">
      <h3>{{title}}</h3>
      <span>{{text}}</span>
      <div class="modal-actions">
        <button @click.stop="fCancel" class="btn-secondary">Отмена</button>
        <button v-if="visibilityAccess" @click.stop="fAccess" class="btn-primary">Подтвердить</button>
        <button v-if="visibilityDelete" @click.stop="fDelete" class="btn-danger">Удалить</button>
      </div>
    </div>
  </div>
</template>

<script>
import {useThemeStore} from "@/stores/theme.js";

export default {
  name: "modalConfirmationWindow",
  props:{
    title:{type: String, default:"подтверждение"},
    text:{type: String, default: "Вы уверенны"},
    visibilityAccess:{type: Boolean, default: false},
    visibilityDelete:{type: Boolean, default: false},
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  methods:{
    fCancel(){
      this.$emit("fCancel")
    },
    fAccess(){
      this.$emit("fAccess")
    },
    fDelete(){
      this.$emit("fDelete")
    },
  }
}
</script>

<style scoped>
/* Стили модального окна */
.modal-overlay {
  background-color: rgba(0, 0, 0, 0.5);
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
}

.modal-content {
  padding: 20px;
  border-radius: 6px;
  width: 400px;
  text-align: center;
  border: 1px solid #4b4b4b;

  h3 {
    margin-top: 0;
  }
  p {
    font-size: 14px;
    line-height: 1.5;
  }
}

.modal-actions {
  display: flex;
  justify-content: space-around;
  margin-top: 20px;

  button {
    padding: 8px 16px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
  }

  .btn-secondary {
    background-color: #757575;
    color: white;
    &:hover {
      background-color: #616161;
    }
  }

  .btn-danger {
    background-color: #d32f2f;
    color: white;
    &:hover {
      background-color: #c62828;
    }
  }

  .btn-primary {
    background-color: #1976d2;
    color: white;
    &:hover {
      background-color: #1565c0;
    }
  }
}
</style>