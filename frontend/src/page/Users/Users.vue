<template>
  <div class="profile-page">
    <div class="table-container" :style="{backgroundColor: themeStore.getColor3, color: themeStore.getColorText}">
      <my-table
          :headerTable="headers"
          :bodyTable="profiles"
      >
        <!-- Слот для колонки логина -->
        <template #login="{ row }">
          <span>{{row.login}}</span>
        </template>

        <!-- Слот для колонки уровня доступа -->
        <!-- Добавили rowKey в деструктуризацию -->
        <template #accessLevel="{ row, rowKey }">

          <select-base
              :value="row.accessLevel"
              :options="accessStatus"
              class="select-field"
              @change="requestAccessChange(rowKey, row, $event)"
          />
        </template>

        <!-- Слот для колонки действий (удаление) -->
        <template #actions="{ rowKey, row }">
          <button-base
              @click="requestDelete(rowKey, row)"
              class="btn-delete"
              value="Удалить"
          />
        </template>
      </my-table>
    </div>

    <!-- Модальное окно подтверждения удаления -->
    <modal-confirmation-window
        v-if="profileToDelete"
        title="Подтверждение удаления"
        :text="`Вы действительно хотите удалить профиль ${profileToDelete.login}?`"
        :visibilityDelete=true
        @fCancel="cancelDelete"
        @fDelete="confirmDelete"
    />

    <!-- Модальное окно подтверждения изменения уровня доступа -->
    <modal-confirmation-window
        v-if="accessChangePending"
        title="Подтверждение удаления"
        :text="`Изменить уровень доступа пользователя  ${accessChangePending.login}? \nс «${ getRoleName(accessChangePending.oldTitle) }» на «${ getRoleName(accessChangePending.newTitle) }»?`"
        :visibilityAccess=true
        @fCancel="cancelAccessChange"
        @fAccess="confirmAccessChange"
    />
  </div>
</template>

<script>
import MyTable from "@/components/UI/MyTable.vue";
import {useThemeStore} from "@/stores/theme.js";
import SelectBase from "@/components/UI/selectBase.vue";
import ButtonBase from "@/components/UI/ButtonBase.vue";
import ModalConfirmationWindow from "@/components/modalWindow/ModalСonfirmationWindow.vue";

export default {
  name: "ProfilePage",
  components: {
    ModalConfirmationWindow,
    ButtonBase,
    SelectBase,
    myTable: MyTable
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  data() {
    return {
      titleHeader: "редактирование пользователей",
      accessStatus: [{value:2, title:"админ"},{value:1, title:"пользователь"}],
      headers: [
        { key: 'login', value: 'Логин', width: '2fr' },
        { key: 'accessLevel', value: 'Уровень доступа', width: '2fr' },
        { key: 'actions', value: 'Действия', width: '1fr' }
      ],
      profiles: {
        'id-1': { login: 'user', accessLevel: 1 },
        'id-2': { login: 'admin', accessLevel: 2 },
        'id-3': { login: 'test', accessLevel: 1 }
      },
      profileToDelete: null,
      // Хранение данных для подтверждения смены роли
      accessChangePending: null
    };
  },
  methods: {
    // Вспомогательный метод для перевода роли на русский язык в диалоге
    getRoleName(role) {
      const roles = {
        'user': 'Пользователь',
        'admin': 'Администратор'
      };
      return roles[role] || role;
    },

    // --- Логика удаления ---
    requestDelete(key, profile) {
      this.profileToDelete = {
        key: key,
        login: profile.login
      };
    },
    cancelDelete() {
      this.profileToDelete = null;
    },
    confirmDelete() {
      if (this.profileToDelete) {
        delete this.profiles[this.profileToDelete.key];
        this.profileToDelete = null;
      }
    },

    // --- Логика изменения уровня доступа ---
    requestAccessChange(key, row, event) {
      const newValue = parseInt(event.target.value);
      const oldValue = row.accessLevel;

      // Если значение не изменилось, ничего не делаем
      if (newValue === oldValue) return;
      this.accessChangePending = {
        key: key,
        login: row.login,
        oldValue: oldValue,
        newValue: newValue,
        oldTitle: this.accessStatus.find(opt => opt.value === oldValue).title,
        newTitle: this.accessStatus.find(opt => opt.value === newValue).title,
        domElement: event.target // сохраняем ссылку на select, чтобы вернуть значение назад при отмене
      };
    },
    cancelAccessChange() {
      if (this.accessChangePending) {
        // Возвращаем визуальный выбор в селекте к исходному значению
        this.accessChangePending.domElement.value = this.accessChangePending.oldValue;
        this.accessChangePending = null;
      }
    },
    confirmAccessChange() {
      if (this.accessChangePending) {
        const { key, newValue } = this.accessChangePending;
        // Применяем новое значение к объекту профиля
        this.profiles[key].accessLevel = newValue;
        this.accessChangePending = null;
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
.profile-page {
  width: 100%;
  height: calc(100% - 110px);
  padding: 50px;
  box-sizing: border-box;

  .table-container {
    border: 1px solid #4b4b4b;
    height: 100%;
    border-radius: 20px;
    padding: 50px;
    box-sizing: border-box;
    position: relative;
  }

  .select-field {
    width: 100%;
    Height: 30px;
  }

  .btn-delete {
    Height: 30px;
  }


}
</style>