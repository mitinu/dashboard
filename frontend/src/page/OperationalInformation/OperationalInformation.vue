<template>
  <div class="operationalInformation">
    <citizensAppeal
        :dateRange="dataAppeal.dateRange"
        :appeals="dataAppeal.appeals"
    />
    <performingDiscipline
        :dateRange="dataAssignmentTest.dateRange"
        :assignments="dataAssignmentTest.assignments"
        :lengths="dataAssignmentTest.lengths"
    />
  </div>
  </template>

<script>
import citizensAppeal from "@/page/OperationalInformation/component/CitizensAppeal.vue";
import performingDiscipline from "@/page/OperationalInformation/component/PerformingDiscipline.vue";
export default {
  name: "operationalInformation",
  components: {
    citizensAppeal,
    performingDiscipline
  },
  data(){
    return{
      title: "оперативные сведения",

      dataAppealTest: {
        dateRange: {
          start: "2026.01.01",
          end: "2026.03.01"
        },
        appeals: {
          length: 0,
          roads: {
            key: "roads",
            title: "Дороги",
            value: 580
          },
          ministryHealth: {
            key: "ministryHealth",
            title: "Министерство здравоохранения",
            value: 130
          },
          KeyIsTag5: {
            key: "KeyIsTag5",
            title: "Тег 5",
            value: 30
          },
          HCS: {
            key: "HCS",
            title: "ЖКХ",
            value: 210
          },
          KeyIsTag4: {
            key: "KeyIsTag4",
            title: "Тег 4",
            value: 100
          }
        },
      },
      dataAppeal:{},

      dataAssignmentTest:{
        dateRange: {
          start: "2026.01.01",
          end: "2026.03.01"
        },
        lengths:{
          all: 0,
          completed: 0,
          overdue: 0,
          inProgress: 0,
        },
        assignments: {
          1: {
            id: 1,
            theOfficial: "Анисимов Д.А.",
            all: 24,
            completed: 18,
            overdue: 2,
            inProgress: 4,
          },
          2: {
            id: 2,
            theOfficial: "Белова Е.В.",
            all: 31,
            completed: 25,
            overdue: 1,
            inProgress: 5,
          },
          3: {
            id: 3,
            theOfficial: "Волков С.Н.",
            all: 42,
            completed: 30,
            overdue: 8,
            inProgress: 4,
          },
          4: {
            id: 4,
            theOfficial: "Григорьева М.И.",
            all: 18,
            completed: 15,
            overdue: 0,
            inProgress: 3,
          },
          5: {
            id: 5,
            theOfficial: "Дмитриев А.П.",
            all: 37,
            completed: 28,
            overdue: 5,
            inProgress: 4,
          }
        }
      },
      dataAssignment:{}

    }
  },
  created() {
    this.dataAppeal = this.dataAppealTest
    for (const appealKey in this.dataAppeal.appeals) {
      if (appealKey!="length"){
        this.dataAppeal.appeals.length+=this.dataAppeal.appeals[appealKey].value
      }
    }


    this.dataAssignment = this.dataAssignmentTest
    for (const assignmentKey in this.dataAssignment.assignments){
      const assignment = this.dataAssignment.assignments[assignmentKey]
      this.dataAssignment.lengths.all += assignment.all
      this.dataAssignment.lengths.completed += assignment.completed
      this.dataAssignment.lengths.overdue += assignment.overdue
      this.dataAssignment.lengths.inProgress += assignment.inProgress
    }
  },
  mounted() {
    this.$emit("setTitleMain", this.title)
    this.$emit("setVisibilityButtonReturnMain", false)
  },
}
</script>

<style scoped>
.operationalInformation{
  height: calc(100% - 110px);
  display: grid;
  grid-template-columns: 2fr 3fr;
  gap: 50px;
  padding: 0 50px;

}
</style>