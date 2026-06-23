<template>
  <div class="finance">
    <div class="content industrialProduction" :style="{backgroundColor: themeStore.getColor3}">
      <chartsPage
          title="промышленного производства"
          :dataset="datasetIndustrialProduction"
          :dates="dates"
          :legend="legendIndustrialProduction"
      />
    </div>
    <div class="content RetailTradeTurnover" :style="{backgroundColor: themeStore.getColor3}">
      <chartsPage
        title="Оборот розничной торговли"
        :dataset="datasetRetailTradeTurnover"
        :dates="dates"
        :legend="legendRetailTradeTurnover"
      />
    </div>
    <div class="content Construction" :style="{backgroundColor: themeStore.getColor3}">
      <chartsPage
        title="Динамика строительных работ"
        :dataset="datasetConstruction"
        :dates="dates"
        :legend="legendConstruction"
      />
    </div>
  </div>
</template>

<script>
import {useThemeStore} from "@/stores/theme.js";
import histogramAndChartPage from "@/components/diagram/page/HistogramAndChartPage.vue";
import ComparisonHistogramPage from "@/components/diagram/page/comparisonHistogramPage.vue";
import chartsPage from "@/components/diagram/page/ChartsPage.vue";

export default {
  name: "finance",
  components:{
    chartsPage,
    ComparisonHistogramPage,
    histogramAndChartPage
  },
  setup() {
    const themeStore = useThemeStore();
    return { themeStore };
  },
  data(){
    return{
      title: "Финансы",
      titleHeader: "Сводные данные",

      datesTest: [
        "4к. 26",
        "3к. 26",
        "2к. 26",
        "1к. 26",
        "4к. 25",
        "3к. 25",
        "2к. 25",
        "1к. 25",
        "4к. 24",
        "3к. 24",
        "2к. 24",
        "1к. 24",
        "4к. 23",
        "3к. 23",
        "2к. 23",
        "1к. 23"
      ],
      dates: [],

      datasetIndustrialProductionTest:{
        chartsColor: [
          {
            "title": "Водоснабжение, водоотведение, организация сбора и утилизация загрязнений",
            "color": "#b49cd0"
          },
          {
            "title": "Обеспечение электрической энергией, газом и паром; кондиционирование воздуха",
            "color": "#6ba5c0"
          },
          {
            "title": "Обрабатывающая промышленность",
            "color": "#e8a735"
          },
          {
            "title": "Добывающая промышленность",
            "color": "#b62c2c"
          }
        ],
        charts: [
          // Водоснабжение, водоотведение, организация сбора и утилизация загрязнений
          {
            "1к. 23": 63.1,
            "2к. 23": 72.5,
            "3к. 23": 82.0,
            "4к. 23": 91.3,
            "1к. 24": 100.7,
            "2к. 24": 103.2,
            "3к. 24": 105.8,
            "4к. 24": 107.5,
            "1к. 25": 108.8,
            "2к. 25": 110.2,
            "3к. 25": 111.6,
            "4к. 25": 113.0,
            "1к. 26": 114.5,
            "2к. 26": 115.9,
            "3к. 26": 117.3,
            "4к. 26": 118.8
          },
          // Обеспечение электрической энергией, газом и паром; кондиционирование воздуха
          {
            "1к. 23": 98.5,
            "2к. 23": 101.2,
            "3к. 23": 103.8,
            "4к. 23": 105.5,
            "1к. 24": 106.8,
            "2к. 24": 110.0,
            "3к. 24": 114.5,
            "4к. 24": 120.3,
            "1к. 25": 127.8,
            "2к. 25": 130.5,
            "3к. 25": 133.2,
            "4к. 25": 136.0,
            "1к. 26": 138.9,
            "2к. 26": 141.8,
            "3к. 26": 144.7,
            "4к. 26": 147.6
          },
          // Обрабатывающая промышленность
          {
            "1к. 23": 99.9,
            "2к. 23": 104.5,
            "3к. 23": 109.2,
            "4к. 23": 114.8,
            "1к. 24": 122.2,
            "2к. 24": 118.5,
            "3к. 24": 112.3,
            "4к. 24": 105.6,
            "1к. 25": 96.7,
            "2к. 25": 100.2,
            "3к. 25": 103.8,
            "4к. 25": 107.5,
            "1к. 26": 111.2,
            "2к. 26": 115.0,
            "3к. 26": 118.8,
            "4к. 26": 122.6
          },
          // Добывающая промышленность
          {
            "1к. 23": 3.7,
            "2к. 23": 28.5,
            "3к. 23": 56.2,
            "4к. 23": 92.8,
            "1к. 24": 136.5,
            "2к. 24": 118.3,
            "3к. 24": 101.2,
            "4к. 24": 85.6,
            "1к. 25": 71.3,
            "2к. 25": 78.9,
            "3к. 25": 86.5,
            "4к. 25": 94.2,
            "1к. 26": 102.0,
            "2к. 26": 109.8,
            "3к. 26": 117.6,
            "4к. 26": 125.4
          }
        ],
        chartConfig: {},
        chartYAxisTitle: "мл.руб."
      },
      datasetIndustrialProduction:{},
      legendIndustrialProduction:{},

      datasetRetailTradeTurnoverTest: {
        chartsColor: [
          {
            "title": "Оборот розничной торговли",
            "color": "#b49cd0"
          },
          {
            "title": "Оборот общественного питания",
            "color": "#6ba5c0"
          },
          {
            "title": "Объем платных услуг",
            "color": "#e8a735"
          }
        ],
        charts: [
          {
            "1к. 23": 35890,
            "2к. 23": 36800,
            "3к. 23": 37800,
            "4к. 23": 38800,
            "1к. 24": 39800,
            "2к. 24": 40800,
            "3к. 24": 41800,
            "4к. 24": 42800,
            "1к. 25": 43800,
            "2к. 25": 44800,
            "3к. 25": 45800,
            "4к. 25": 46800,
            "1к. 26": 47800,
            "2к. 26": 48800,
            "3к. 26": 49800,
            "4к. 26": 50800
          },
          {
            "1к. 23": 2026,
            "2к. 23": 2150,
            "3к. 23": 2300,
            "4к. 23": 2480,
            "1к. 24": 2650,
            "2к. 24": 2840,
            "3к. 24": 3050,
            "4к. 24": 3280,
            "1к. 25": 3520,
            "2к. 25": 3780,
            "3к. 25": 4060,
            "4к. 25": 4360,
            "1к. 26": 4680,
            "2к. 26": 5020,
            "3к. 26": 5380,
            "4к. 26": 5760
          },
          {
            "1к. 23": 16500,
            "2к. 23": 17200,
            "3к. 23": 18100,
            "4к. 23": 19000,
            "1к. 24": 19800,
            "2к. 24": 20400,
            "3к. 24": 21100,
            "4к. 24": 20623,
            "1к. 25": 25916,
            "2к. 25": 26800,
            "3к. 25": 27700,
            "4к. 25": 28500,
            "1к. 26": 29400,
            "2к. 26": 30300,
            "3к. 26": 31200,
            "4к. 26": 32100
          }
        ],
        chartConfig: {},
        chartYAxisTitle: "мл.руб."
      },
      datasetRetailTradeTurnover:{},
      legendRetailTradeTurnover:{},

      datasetConstructionTest: {
        chartsColor: [
          {
            "title": "Строительные работы, млрд рублей",
            "color": "#4CAF50"
          }
        ],
        chartsColor2: [
          {
            "title": "Ввод жилья, тыс. кв. метров",
            "color": "#FF9800"
          }
        ],
        charts: [
          {
            "1к. 23": 9.2,
            "2к. 23": 9.8,
            "3к. 23": 10.5,
            "4к. 23": 11.3,
            "1к. 24": 7.7,
            "2к. 24": 8.4,
            "3к. 24": 9.1,
            "4к. 24": 10.2,
            "1к. 25": 12.1,
            "2к. 25": 12.8,
            "3к. 25": 13.5,
            "4к. 25": 14.2,
            "1к. 26": 14.9,
            "2к. 26": 15.6,
            "3к. 26": 16.3,
            "4к. 26": 17.0
          }
        ],
        charts2: [
          {
            "1к. 23": 45.5,
            "2к. 23": 47.2,
            "3к. 23": 48.8,
            "4к. 23": 50.1,
            "1к. 24": 50.2,
            "2к. 24": 48.5,
            "3к. 24": 46.8,
            "4к. 24": 44.2,
            "1к. 25": 27.5,
            "2к. 25": 29.1,
            "3к. 25": 30.8,
            "4к. 25": 32.5,
            "1к. 26": 34.2,
            "2к. 26": 36.0,
            "3к. 26": 37.8,
            "4к. 26": 39.5
          }
        ],
        chartConfig: {
          // YAxisTitle: 2
        },
        chartYAxisTitle: "млрд.руб.",
        chartYAxisTitle2: "тыс.кв.м."
      },
      datasetConstruction:{},
      legendConstruction:{},
    }
  },
  created() {
    this.dates = this.datesTest

    this.datasetIndustrialProduction = this.datasetIndustrialProductionTest
    this.legendIndustrialProduction = [this.datasetIndustrialProduction.chartsColor[0], this.datasetIndustrialProduction.chartsColor[1], this.datasetIndustrialProduction.chartsColor[2], this.datasetIndustrialProduction.chartsColor[3]]

    this.datasetRetailTradeTurnover = this.datasetRetailTradeTurnoverTest
    this.legendRetailTradeTurnover = [this.datasetRetailTradeTurnover.chartsColor[0], this.datasetRetailTradeTurnover.chartsColor[1], this.datasetRetailTradeTurnover.chartsColor[2]]

    this.datasetConstruction = this.datasetConstructionTest
    this.legendConstruction = [this.datasetConstruction.chartsColor[0], this.datasetConstruction.chartsColor2[0]]



  },
  mounted() {
    this.$emit("setTitleMain", "")
    this.$emit("setTitleHeader", this.title)
    this.$emit("setVisibilityButtonReturnMain", false)
  },
}
</script>

<style scoped>
.finance{
  height: calc(100% - 110px);
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  padding: 0 50px;
  gap: 50px;
  .industrialProduction {
    grid-column: 1 / 3;
  }
  .content{
    border: 1px solid #4b4b4b;
    border-radius: 20px;
    min-width: 0;
  }
}
</style>
