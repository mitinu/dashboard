<template>
  <div class="citizensAppeal">
    <header-information-block
        title="обращение граждан"
        :dateRange="dateRange"
    />
    <div class="body">
      <div class="container">
        <div><span>сообщение граждан</span></div>
        <div>
          <countMessage
              color="#ffffff"
              :count="appeals.length"
              title="Всего"
          />
        </div>
      </div>
      <div class="container fl1">
        <myTable
            class="myTable"
            :headerTable="headerTable"
            :bodyTable="tableRows"
        />
      </div>
    </div>
  </div>
</template>

<script>
import HeaderInformationBlock from "@/page/OperationalInformation/component/HeaderInformationBlock.vue";
import MessageIcon from "@/components/icons/MessageIcon.vue";
import MyTable from "@/components/UI/MyTable.vue";
import CountMessage from "@/page/OperationalInformation/component/CountMessage.vue";

export default {
  name: "citizensAppeal",
  components: {
    CountMessage,
    MyTable,
    MessageIcon,
    HeaderInformationBlock
  },
  props:{
    dateRange:{type: Object},
    appeals:{type: Object}
  },
  data(){
    return{
      headerTable:[
        {
          key:"number",
          value:"№",
          width:"1fr"
        },
        {
          key:"topic",
          value:"тема",
          width:"2fr"
        },
        {
          key:"count",
          value:"кол-во",
          width:"1fr"
        },
      ],
    }
  },
  computed: {
    /**
     * Преобразует объект обращений (appeals) в отсортированный массив для таблицы.
     * Исключает техническое свойство 'length', сортирует по убыванию количества
     * и добавляет порядковую нумерацию.
     * */
    tableRows() {
      // Проверка на наличие данных, чтобы избежать ошибок при рендеринге
      if (!this.appeals) return [];

      return Object.entries(this.appeals)
          .filter(([key]) => key !== 'length')
          .map(([key, value]) => ({ key, ...value }))
          .sort((a, b) => b.value - a.value)
          .map((item, index) => ({
            id: index + 1,
            number: index + 1,
            count: item.value,
            topic: item.title,
          }));
    }
  },
}
</script>

<style scoped>
.citizensAppeal{
  background-color: #171620;
  border-radius: 5%;
  border: 1px solid #4b4b4b;
  display: flex;
  flex-direction: column;
  .body{
    display: flex;
    justify-content: space-between;
    flex: 1;
    .container{
      padding: 20px;
      padding-bottom: 50px;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      border-right: 1px solid #4b4b4b;
      &:last-child{
        border: none;
        flex: 1;
      }
      .myTable{
        height: 100%;
      }
    }
  }
}
</style>