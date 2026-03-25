<template>
  <div class="citizensAppeal">
    <header-information-block
        title="исполнительская дисциплина"
        :dateRange="dateRange"
    />
    <div class="body">
      <div class="container">
        <div><span>поручение</span></div>
        <div>
          <countMessage
            color="#ffffff"
            :count="lengths.all"
            title="Всего"
          />
          <countMessage
            color="#2cb571"
            :count="lengths.completed"
            title="завершённых"
          />
          <countMessage
            color="royalblue"
            :count="lengths.inProgress"
            title="в процессе"
          />
          <countMessage
            color="#ce3233"
            :count="lengths.overdue"
            title="просроченных"
          />
        </div>
      </div>
      <div class="container fl1">
        <myTable
          class="myTable"
          :headerTable="dataTable.header"
          :bodyTable="dataTable.body"
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
  name: "performingDiscipline",
  components: {
    CountMessage,
    MyTable,
    MessageIcon,
    HeaderInformationBlock
  },
  props:{
    dateRange:{type: Object},
    assignments:{type: Array},
    lengths:{type: Object}
  },
  data(){
    return{
      dataTable:{
        header:[
          {
            key:"theOfficial",
            value:"Глава города",
            width:"2fr"
          },
          {
            key:"all",
            value:"всего",
            width:"1fr"
          },
          {
            key:"completed",
            value:"завершённые",
            width:"1fr"
          },
          {
            key:"inProgress",
            value:"в\u00A0процессе",
            width:"1fr"
          },
          {
            key:"overdue",
            value:"завершённые",
            width:"1fr"
          },
        ],
        body:this.assignments
      }
    }
  },
  mounted() {
  //   Object.entries(this.appeals)
  //     .filter(([key]) => key !== 'length')
  //     .map(([key, value]) => ({ key, ...value }))
  //     .sort((a, b) => b.value - a.value)
  //     .forEach((item, index) => {
  //       this.dataTable.body[index+1] = {
  //         id:index+1,
  //         number:index+1,
  //         count:item.value,
  //         topic:item.title,
  //       }
  //     });
  }
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