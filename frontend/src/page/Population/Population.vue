<template>
  <div class="population">
    <div class="content populationSize" :style="{backgroundColor: themeStore.getColor3}">
      <histogramAndChartPage
          title="Численность население"
          :dataset="datasetPopulation"
          :dates="dates"
          :legend="legendPopulation"
      />
    </div>
    <div class="content naturalMovementPopulation" :style="{backgroundColor: themeStore.getColor3}">
      <comparisonHistogramPage
          title="Натуральное движение"
          :dataset="datasetNaturalMovement"
          :dates="dates"
          :legend="legendNaturalMovement"
      />
    </div>
    <div class="content artificialMovementPopulation" :style="{backgroundColor: themeStore.getColor3}">
      <comparisonHistogramPage
          title="Искусственное движение"
          :dataset="datasetArtificialMovement"
          :dates="dates"
          :legend="legendArtificialMovement"
      />
    </div>
  </div>
</template>

<script>
import {useThemeStore} from "@/stores/theme.js";
import histogramAndChartPage from "@/components/diagram/page/HistogramAndChartPage.vue";
import ComparisonHistogramPage from "@/components/diagram/page/comparisonHistogramPage.vue";

export default {
  name: "population",
  components:{
    ComparisonHistogramPage,
    histogramAndChartPage
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  data(){
    return{
      title: "Движение населения",
      titleHeader: "Сводные данные",

      datesTest: [
        "27.09.25",
        "27.06.25",
        "27.03.25",
        "26.12.25",
        "26.09.25",
        "26.06.25",
        "26.03.25",
        "25.12.25",
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
      dates: [],


      datasetPopulationTest: {
        histogramColor: {
          title: "прирост/убыток населения, чел.",
          color: "#412467"
        },
        histogram: {
          "22.09.25": 320,
          "22.12.25": 580,
          "23.03.25": 410,
          "23.06.25": -240,
          "23.09.25": 280,
          "23.12.25": 650,
          "24.03.25": 390,
          "24.06.25": -180,
          "24.09.25": 340,
          "24.12.25": 720,
          "25.03.25": 460,
          "25.06.25": -310,
          "25.09.25": 510,
          // === ТЕСТОВЫЕ ДАННЫЕ (продолжение) ===
          "25.12.25": 680,
          "26.03.25": 430,
          "26.06.25": -260,
          "26.09.25": 490,
          "26.12.25": 750,
          "27.03.25": 520,
          "27.06.25": -290,
          "27.09.25": 560,
        },
        histogramConfig: {},
        histogramYAxisTitle: "чел",
        chartColor: {
          title: "численность населения, чел.",
          color: "#b49cd0"
        },
        chart: {
          "22.09.25": 548200,
          "22.12.25": 548780,
          "23.03.25": 549190,
          "23.06.25": 548950,
          "23.09.25": 549230,
          "23.12.25": 549880,
          "24.03.25": 550270,
          "24.06.25": 550090,
          "24.09.25": 550430,
          "24.12.25": 551150,
          "25.03.25": 551610,
          "25.06.25": 551300,
          "25.09.25": 551810,
          // === ТЕСТОВЫЕ ДАННЫЕ (продолжение) ===
          "25.12.25": 552490,  // +680
          "26.03.25": 552920,  // +430
          "26.06.25": 552660,  // -260
          "26.09.25": 553150,  // +490
          "26.12.25": 553900,  // +750
          "27.03.25": 554420,  // +520
          "27.06.25": 554130,  // -290
          "27.09.25": 554690,  // +560
        },
        chartConfig: {},
        chartYAxisTitle: "человек"
      },

      datasetPopulation: {},
      legendPopulation: [],

      datasetNaturalMovementTest: {
        histogramColor: [
          {
            title: "прирост населения",
            color: "#412467"
          },
          {
            title: "убыток населения",
            color: "#b49cd0"
          },
        ],
        histogram: {
          "22.09.25": [520, 340],   // [рождаемость, смертность] = +180 естественный прирост
          "22.12.25": [680, 330],   // = +350
          "23.03.25": [580, 340],   // = +240
          "23.06.25": [420, 540],   // = -120 (смертность выше рождаемости)
          "23.09.25": [500, 340],   // = +160
          "23.12.25": [720, 330],   // = +390
          "24.03.25": [560, 330],   // = +230
          "24.06.25": [430, 520],   // = -90
          "24.09.25": [530, 330],   // = +200
          "24.12.25": [760, 330],   // = +430
          "25.03.25": [600, 330],   // = +270
          "25.06.25": [410, 570],   // = -160
          "25.09.25": [630, 330],   // = +300
          "25.12.25": [730, 330],   // = +400
          "26.03.25": [580, 330],   // = +250
          "26.06.25": [440, 570],   // = -130
          "26.09.25": [620, 330],   // = +290
          "26.12.25": [770, 330],   // = +440
          "27.03.25": [640, 330],   // = +310
          "27.06.25": [430, 580],   // = -150
          "27.09.25": [660, 330],   // = +560
        },
        histogramConfig: {},
        histogramYAxisTitle: "человек",

      },

      datasetNaturalMovement: {},
      legendNaturalMovement: [],

      datasetArtificialMovementTest: {
        histogramColor: [
          {
            title: "прибытие населения",
            color: "#412467"
          },
          {
            title: "убыток населения",
            color: "#b49cd0"
          },
        ],
        histogram: {
          "22.09.25": [420, 280],   // [прибытие, выбытие] = +140 миграционный прирост
          "22.12.25": [510, 280],   // = +230
          "23.03.25": [460, 290],   // = +170
          "23.06.25": [340, 460],   // = -120 (отток выше прибытия)
          "23.09.25": [410, 290],   // = +120
          "23.12.25": [540, 280],   // = +260
          "24.03.25": [450, 290],   // = +160
          "24.06.25": [360, 450],   // = -90
          "24.09.25": [430, 290],   // = +140
          "24.12.25": [570, 280],   // = +290
          "25.03.25": [480, 290],   // = +190
          "25.06.25": [330, 480],   // = -150
          "25.09.25": [500, 290],   // = +210
          "25.12.25": [560, 280],   // = +280
          "26.03.25": [470, 290],   // = +180
          "26.06.25": [350, 480],   // = -130
          "26.09.25": [490, 290],   // = +200
          "26.12.25": [590, 280],   // = +310
          "27.03.25": [500, 290],   // = +210
          "27.06.25": [340, 490],   // = -150
          "27.09.25": [520, 290],   // = +230
        },
        histogramConfig: {},
        histogramYAxisTitle: "человек"
      },

      datasetArtificialMovement: {},
      legendArtificialMovement: [],
    }
  },
  created() {
    this.dates = this.datesTest

    this.datasetPopulation = this.datasetPopulationTest
    this.legendPopulation = [this.datasetPopulation.histogramColor, this.datasetPopulation.chartColor]

    this.datasetNaturalMovement = this.datasetNaturalMovementTest
    this.legendNaturalMovement = [this.datasetNaturalMovement.histogramColor[0], this.datasetNaturalMovement.histogramColor[1]]

    this.datasetArtificialMovement = this.datasetArtificialMovementTest
    this.legendArtificialMovement = [this.datasetArtificialMovement.histogramColor[0], this.datasetArtificialMovement.histogramColor[1]]

  },
  mounted() {
    this.$emit("setTitleMain", "")
    this.$emit("setTitleHeader", this.title)
    this.$emit("setVisibilityButtonReturnMain", false)
  },
}
</script>

<style scoped>
.population{
  height: calc(100% - 110px);
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  padding: 0 50px;
  gap: 50px;
  .populationSize {
    grid-column: 1 / 3;
  }
  .content{
    border: 1px solid #4b4b4b;
    border-radius: 20px;
    min-width: 0;
  }
}
</style>
