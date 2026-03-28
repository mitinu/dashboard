<template>
  <div class="workingCapacityPopulation">
    <div class="content unemployment">
      <histogramAndChartPage
          title="уровень безработицы"
          :dataset="datasetUnemployment"
          :dates="datesUnemployment"
          :legend="legendUnemployment"
      />
    </div>
    <div class="content averageSalary">
      <chartPage
          title="среднемесячная начисляемая заработная плата"
          :dataset="datasetAverageSalary"
          :dates="datesAverageSalary"
          :legend="legendAverageSalary"
      />
    </div>
    <div class="content CPI">
      <fourChartPage
          title="индекс потребительских цен к соответствующим кварталам прошлых лет, %"
          :dataset="datasetCPI"
          :dates="datesCPI"
          :legend="legendCPI"
      />
    </div>
  </div>
</template>

<script>
import histogramAndChartPage from "@/components/diagram/page/HistogramAndChartPage.vue";
import chartPage from "@/components/diagram/page/chartPage.vue";
import fourChartPage from "@/components/diagram/page/FourChartPage.vue";

export default {
  components:{
    histogramAndChartPage,
    chartPage,
    fourChartPage
  },
  name: "monthlyInformation",
  data(){
    return{
      title: "Ежемесячная информация",

      datasetUnemploymentTest: {
        histogramColor:{
          title: "численность безработных, чел.",
          color: "#412467"
        },
        histogram: {
          "22.09.25": 1726,
          "22.12.25": 1541,
          "23.03.25": 1421,
          "23.06.25": 1304,
          "23.09.25": 1166,
          "23.12.25": 1026,
          "24.03.25": 946,
          "24.06.25": 817,
          "24.09.25": 746,
          "24.12.25": 752,
          "25.03.25": 792,
          "25.06.25": 829,
          "25.09.25": 1064,
        },
        histogramConfig:{},
        histogramYAxisTitle:"чел",
        chartColor:{
          title: "уровень безработицы %",
          color: "#b49cd0"
        },
        chart: {
          "22.09.25": 0.53,
          "22.12.25": 0.45,
          "23.03.25": 0.42,
          "23.06.25": 0.39,
          "23.09.25": 0.36,
          "23.12.25": 0.36,
          "24.03.25": 0.29,
          "24.06.25": 0.25,
          "24.09.25": 0.23,
          "24.12.25": 0.23,
          "25.03.25": 0.24,
          "25.06.25": 0.25,
          "25.09.25": 0.32,
        },
        chartConfig:{},
        chartYAxisTitle:"%"

      },
      datesUnemploymentTest: [
        "25.09.25",
        "25.06.25",
        "25.03.25",
        "24.12.25",
        "24.09.25",
        "24.06.25",
        "24.03.25",
        "23.12.25",
        "23.09.25",
        "23.06.25",
        "23.03.25",
        "22.12.25",
        "22.09.25",
      ],
      datasetUnemployment: {},
      datesUnemployment: [],
      legendUnemployment: [],

      datasetAverageSalaryTest: {
        chartColor:{
          title: "уровень безработицы %",
          color: "#b49cd0"
        },
        chart: {
          "4к. 25": 89000,
          "3к. 25": 84000,
          "2к. 25": 80000,
          "1к. 25": 92000,
        },
        chartConfig:{
          lowerBackgroundFill: true
        },
      },
      datesAverageSalaryTest: [
        "4к. 25",
        "3к. 25",
        "2к. 25",
        "1к. 25",
      ],
      datasetAverageSalary: {},
      datesAverageSalary: [],
      legendAverageSalary: [],

      datasetCPITest: {
        chart1Color:{
          title: "2022",
          color: "#b49cd0",
        },
        chart1: {
          "сен 22": 112,
          "дек 22": 108,
          "мар 23": 105,
          "июн 23": 102,
          "сен 23": 101,
          "дек 23": 100,
          "мар 24": 104,
          "июн 24": 107,
          "сен 24": 110,
          "дек 24": 115,
          "мар 25": 120,
          "июн 25": 125,
          "сен 25": 130,
        },
        chart2Color:{
          title: "2023",
          color: "#6ba5c0",
        },
        chart2: {
          "сен 22": 145,
          "дек 22": 142,
          "мар 23": 138,
          "июн 23": 135,
          "сен 23": 132,
          "дек 23": 130,
          "мар 24": 128,
          "июн 24": 132,
          "сен 24": 138,
          "дек 24": 144,
          "мар 25": 150,
          "июн 25": 155,
          "сен 25": 160,
        },
        chart3Color:{
          title: "2024",
          color: "#e8a735",
        },
        chart3: {
          "сен 22": 128,
          "дек 22": 125,
          "мар 23": 122,
          "июн 23": 118,
          "сен 23": 115,
          "дек 23": 112,
          "мар 24": 118,
          "июн 24": 125,
          "сен 24": 132,
          "дек 24": 138,
          "мар 25": 142,
          "июн 25": 148,
          "сен 25": 152,
        },
        chart4Color:{
          title: "2025",
          color: "#e87c6b",
        },
        chart4: {
          "сен 22": 105,
          "дек 22": 110,
          "мар 23": 115,
          "июн 23": 120,
          "сен 23": 125,
          "дек 23": 130,
          "мар 24": 135,
          "июн 24": 140,
          "сен 24": 152,
          "дек 24": 145,
          "мар 25": 150,
          "июн 25": 155,
          "сен 25": 158,
        },
        chartConfig:{},
      },
      datesCPITest: [
        "сен 25",
        "июн 25",
        "мар 25",
        "дек 24",
        "сен 24",
        "июн 24",
        "мар 24",
        "дек 23",
        "сен 23",
        "июн 23",
        "мар 23",
        "дек 22",
        "сен 22",
      ],
      datasetCPI: {},
      datesCPI: [],
      legendCPI: []

    }
  },
  created() {
    this.datasetUnemployment = this.datasetUnemploymentTest
    this.datesUnemployment = this.datesUnemploymentTest
    this.legendUnemployment = [this.datasetUnemployment.histogramColor, this.datasetUnemployment.chartColor]

    this.datasetAverageSalary = this.datasetAverageSalaryTest
    this.datesAverageSalary = this.datesAverageSalaryTest

    this.datasetCPI = this.datasetCPITest
    this.datesCPI = this.datesCPITest
    this.legendCPI = [this.datasetCPI.chart1Color, this.datasetCPI.chart2Color, this.datasetCPI.chart3Color, this.datasetCPI.chart4Color]
  },
  mounted() {
    this.$emit("setTitleMain", this.title)
    this.$emit("setVisibilityButtonReturnMain", false)
  },

}
</script>

<style scoped>
.workingCapacityPopulation{
  height: calc(100% - 110px);
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  padding: 0 50px;
  gap: 50px;
  .CPI {
    grid-column: 1 / 3;
  }
  .content{
    border: 1px solid #4b4b4b;
    background-color: #171620;
    border-radius: 20px;
    min-width: 0;
  }
}
</style>