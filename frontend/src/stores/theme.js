import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useThemeStore = defineStore('theme', () => {
    // state (данные)
    const backgroundMain = ref("#f7f9ff")
    const colorText = ref("#0b0b10")
    const color3 = ref("#E8E6F5")
    const color4 = ref("#D0D0E0")


    // getters (вычисляемые свойства)
    const getBackgroundMain = computed(() => backgroundMain.value)
    const getColorText = computed(() => colorText.value)
    const getColor3 = computed(() => color3.value)
    const getColor4 = computed(() => color4.value)

    // actions (методы для изменения данных)
    function setBackgroundMain(newBackgroundMain) {
        backgroundMain.value = newBackgroundMain
    }
    function setColorText(newColorText) {
        colorText.value = newColorText
    }
    function setColor3(newColor) {
        color3.value = newColor
    }
    function setColor4(newColor) {
        color4.value = newColor
    }

    return { getBackgroundMain, setBackgroundMain, getColorText, setColorText, getColor3, setColor3, getColor4, setColor4}
})