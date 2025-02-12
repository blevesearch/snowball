package russian

import (
	"github.com/blevesearch/snowball/romance"
	"testing"
)

// Test isLowerVowel for things we know should be true
// or false.
//
func Test_isLowerVowel(t *testing.T) {
	testCases := []romance.WordBoolTestCase{
		// These are all vowels.
		{"аеиоуыэюя", true},
		// None of these are vowels.
		{"бвгджзйклмнпрстфхцчшщъь", false},
	}
	romance.RunRunewiseBoolTest(t, isLowerVowel, testCases)
}

// Test stopWords for things we know should be true
// or false.
//
func Test_stopWords(t *testing.T) {
	testCases := []romance.WordBoolTestCase{
		{"была", true},
		{"нас", true},
		{"меня", true},
		{"химическое", false},
		{"машиностроение", false},
	}
	romance.RunWordBoolTest(t, isStopWord, testCases)
}

func Test_findRegions(t *testing.T) {
	testCases := []romance.FindRegionsTestCase{
		{"кистей", 3, 6, 2},
		{"пугает", 3, 6, 2},
		{"горячей", 3, 5, 2},
		{"разочароваться", 3, 5, 2},
		{"ваших", 3, 5, 2},
		{"сливаешься", 4, 7, 3},
		{"сбавки", 4, 6, 3},
		{"выпил", 3, 5, 2},
		{"фирс", 3, 4, 2},
		{"ездит", 2, 5, 1},
		{"думай", 3, 5, 2},
		{"летай", 3, 5, 2},
		{"беседовала", 3, 5, 2},
		{"беседой", 3, 5, 2},
		{"брюзжит", 4, 7, 3},
		{"растущего", 3, 6, 2},
		{"подползла", 3, 6, 2},
		{"быстротою", 3, 7, 2},
		{"портит", 3, 6, 2},
		{"качала", 3, 5, 2},
		{"ободрившийся", 2, 4, 1},
		{"чуткость", 3, 6, 2},
		{"полусне", 3, 5, 2},
		{"стремление", 5, 8, 4},
		{"рубинштейн", 3, 5, 2},
		{"прихоти", 4, 6, 3},
		{"чувстве", 3, 7, 2},
		{"существе", 3, 5, 2},
		{"горстью", 3, 7, 2},
		{"постелью", 3, 6, 2},
		{"приостановился", 5, 8, 3},
		{"бегает", 3, 6, 2},
		{"неприличия", 3, 6, 2},
		{"терзаешь", 3, 7, 2},
		{"лягу", 3, 4, 2},
		{"недоуменьем", 3, 6, 2},
		{"скрасить", 5, 7, 4},
		{"нелепый", 3, 5, 2},
		{"измениться", 2, 5, 1},
		{"бульвару", 3, 7, 2},
		{"засуетятся", 3, 6, 2},
		{"благосостояния", 4, 6, 3},
		{"поляне", 3, 5, 2},
		{"прощенья", 4, 6, 3},
		{"исподтишка", 2, 5, 1},
		{"переселился", 3, 5, 2},
		{"отмахивается", 2, 5, 1},
		{"измучил", 2, 5, 1},
		{"пьяной", 4, 6, 3},
		{"куртке", 3, 6, 2},
		{"буйный", 3, 6, 2},
		{"эско", 2, 4, 1},
		{"застрахованы", 3, 7, 2},
		{"шаршавого", 3, 6, 2},
		{"ра", 2, 2, 2},
		{"засветло", 3, 6, 2},
		{"накануне", 3, 5, 2},
		{"периодического", 3, 6, 2},
		{"молокососы", 3, 5, 2},
		{"выгода", 3, 5, 2},
		{"всунуть", 4, 6, 3},
		{"мачеха", 3, 5, 2},
		{"ложечки", 3, 5, 2},
		{"счастьи", 4, 7, 3},
		{"дворовый", 4, 6, 3},
		{"замерло", 3, 5, 2},
		{"окружающие", 2, 5, 1},
		{"ят", 2, 2, 1},
		{"обычных", 2, 4, 1},
		{"добежав", 3, 5, 2},
		{"грезилось", 4, 6, 3},
		{"татарским", 3, 5, 2},
		{"неважным", 3, 5, 2},
		{"бежавших", 3, 5, 2},
		{"завести", 3, 5, 2},
		{"отворенную", 2, 5, 1},
		{"праотцам", 5, 8, 3},
		{"дракон", 4, 6, 3},
		{"бродит", 4, 6, 3},
		{"грусть", 4, 6, 3},
		{"постареть", 3, 6, 2},
		{"спасенья", 4, 6, 3},
		{"предыдущая", 4, 6, 3},
		{"расчеты", 3, 6, 2},
		{"плавает", 4, 7, 3},
		{"наступили", 3, 6, 2},
		{"повели", 3, 5, 2},
		{"подчинились", 3, 6, 2},
		{"милостисдарь", 3, 5, 2},
		{"здравый", 5, 7, 4},
		{"представляла", 4, 8, 3},
		{"тронь", 4, 5, 3},
		{"поклонись", 3, 6, 2},
		{"предчувствию", 4, 7, 3},
		{"пропускающим", 4, 6, 3},
		{"щетку", 3, 5, 2},
		{"отворотит", 2, 5, 1},
		{"богом", 3, 5, 2},
		{"кишки", 3, 5, 2},
		{"полярный", 3, 5, 2},
	}

	romance.RunFindRegionsTest(t, findRegions, testCases)
}

// Test step1
//
func Test_step1(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"оснеженные", 2, 5, 1, true, "оснеженн", 2, 5, 1},
		{"умираю", 2, 4, 1, true, "умира", 2, 4, 1},
		{"старосте", 4, 6, 3, true, "старост", 4, 6, 3},
		{"вспугнуть", 5, 8, 4, true, "вспугнут", 5, 8, 4},
		{"отстаньте", 2, 6, 1, true, "отстаньт", 2, 6, 1},
		{"летние", 3, 6, 2, true, "летн", 3, 4, 2},
		{"густой", 3, 6, 2, true, "густ", 3, 4, 2},
		{"кутью", 3, 5, 2, true, "куть", 3, 4, 2},
		{"социальный", 3, 6, 2, true, "социальн", 3, 6, 2},
		{"сшедший", 4, 7, 3, true, "сшедш", 4, 5, 3},
		{"силою", 3, 5, 2, true, "сил", 3, 3, 2},
		{"мельницы", 3, 7, 2, true, "мельниц", 3, 7, 2},
		{"блюдечками", 4, 6, 3, true, "блюдечк", 4, 6, 3},
		{"частном", 3, 7, 2, true, "частн", 3, 5, 2},
		{"сухарь", 3, 5, 2, true, "сухар", 3, 5, 2},
		{"шатаясь", 3, 6, 2, true, "шат", 3, 3, 2},
		{"пылают", 3, 6, 2, true, "пыла", 3, 4, 2},
		{"оглядывает", 2, 5, 1, true, "оглядыва", 2, 5, 1},
		{"многоуважаемая", 4, 7, 3, true, "многоуважа", 4, 7, 3},
		{"носками", 3, 6, 2, true, "носк", 3, 4, 2},
		{"согрешил", 3, 6, 2, true, "согреш", 3, 6, 2},
		{"почесть", 3, 5, 2, true, "почест", 3, 5, 2},
		{"резвому", 3, 6, 2, true, "резв", 3, 4, 2},
		{"сверкнувший", 4, 8, 3, true, "сверкнувш", 4, 8, 3},
		{"стерегла", 4, 6, 3, true, "стерегл", 4, 6, 3},
		{"удалюсь", 2, 4, 1, true, "удал", 2, 4, 1},
		{"бульдоге", 3, 7, 2, true, "бульдог", 3, 7, 2},
		{"житья", 3, 5, 2, true, "жит", 3, 3, 2},
		{"учился", 2, 4, 1, true, "уч", 2, 2, 1},
		{"невысокому", 3, 5, 2, true, "невысок", 3, 5, 2},
		{"сновать", 4, 6, 3, true, "снова", 4, 5, 3},
		{"оли", 2, 3, 1, true, "ол", 2, 2, 1},
		{"освежить", 2, 5, 1, true, "освеж", 2, 5, 1},
		{"хлынуло", 4, 6, 3, true, "хлынул", 4, 6, 3},
		{"ветхий", 3, 6, 2, true, "ветх", 3, 4, 2},
		{"покончив", 3, 5, 2, true, "поконч", 3, 5, 2},
		{"веселее", 3, 5, 2, true, "весел", 3, 5, 2},
		{"пантелеевна", 3, 6, 2, true, "пантелеевн", 3, 6, 2},
		{"трепещущий", 4, 6, 3, true, "трепещущ", 4, 6, 3},
		{"опустошаются", 2, 4, 1, true, "опустоша", 2, 4, 1},
		{"скажите", 4, 6, 3, true, "скаж", 4, 4, 3},
		{"пригрозили", 4, 7, 3, true, "пригроз", 4, 7, 3},
		{"кузнеца", 3, 6, 2, true, "кузнец", 3, 6, 2},
		{"эстетическое", 2, 5, 1, true, "эстетическ", 2, 5, 1},
		{"проделки", 4, 6, 3, true, "проделк", 4, 6, 3},
		{"согласится", 3, 6, 2, true, "соглас", 3, 6, 2},
		{"скуку", 4, 5, 3, true, "скук", 4, 4, 3},
		{"поданы", 3, 5, 2, true, "пода", 3, 4, 2},
		{"походкой", 3, 5, 2, true, "походк", 3, 5, 2},
		{"поважнее", 3, 5, 2, true, "поважн", 3, 5, 2},
		{"лотереей", 3, 5, 2, true, "лотере", 3, 5, 2},
		{"парнике", 3, 6, 2, true, "парник", 3, 6, 2},
		{"содом", 3, 5, 2, true, "сод", 3, 3, 2},
		{"полученному", 3, 5, 2, true, "полученн", 3, 5, 2},
		{"проезжавшую", 5, 8, 3, true, "проезжа", 5, 7, 3},
		{"едучи", 2, 4, 1, true, "едуч", 2, 4, 1},
		{"помнишь", 3, 6, 2, true, "помн", 3, 4, 2},
		{"греховные", 4, 6, 3, true, "греховн", 4, 6, 3},
		{"кухарке", 3, 5, 2, true, "кухарк", 3, 5, 2},
		{"запретишь", 3, 6, 2, true, "запрет", 3, 6, 2},
		{"похолодели", 3, 5, 2, true, "похолодел", 3, 5, 2},
		{"опамятовавшись", 2, 4, 1, true, "опамятова", 2, 4, 1},
		{"валяются", 3, 6, 2, true, "валя", 3, 4, 2},
		{"николаевском", 3, 5, 2, true, "николаевск", 3, 5, 2},
		{"физиология", 3, 6, 2, true, "физиолог", 3, 6, 2},
		{"захромала", 3, 6, 2, true, "захрома", 3, 6, 2},
		{"суйся", 3, 5, 2, true, "су", 2, 2, 2},
		{"бесплодного", 3, 7, 2, true, "бесплодн", 3, 7, 2},
		{"спасения", 4, 6, 3, true, "спасен", 4, 6, 3},
		{"неумолкаемый", 4, 6, 2, true, "неумолка", 4, 6, 2},
		{"идиотством", 2, 5, 1, true, "идиотств", 2, 5, 1},
		{"влюбленной", 4, 7, 3, true, "влюбленн", 4, 7, 3},
		{"родионом", 3, 6, 2, true, "родион", 3, 6, 2},
		{"морщился", 3, 6, 2, true, "морщ", 3, 4, 2},
		{"истерзанной", 2, 5, 1, true, "истерза", 2, 5, 1},
		{"обвинила", 2, 5, 1, true, "обвин", 2, 5, 1},
		{"загадочным", 3, 5, 2, true, "загадочн", 3, 5, 2},
		{"заикнитесь", 4, 7, 2, true, "заикн", 4, 5, 2},
		{"пелериной", 3, 5, 2, true, "пелерин", 3, 5, 2},
		{"ловко", 3, 5, 2, true, "ловк", 3, 4, 2},
		{"обличий", 2, 5, 1, true, "облич", 2, 5, 1},
		{"потасканной", 3, 5, 2, true, "потаска", 3, 5, 2},
		{"ненасытную", 3, 5, 2, true, "ненасытн", 3, 5, 2},
		{"засверкал", 3, 6, 2, true, "засверка", 3, 6, 2},
		{"сидишь", 3, 5, 2, true, "сид", 3, 3, 2},
		{"выступление", 3, 6, 2, true, "выступлен", 3, 6, 2},
		{"дивное", 3, 6, 2, true, "дивн", 3, 4, 2},
		{"распечатывать", 3, 6, 2, true, "распечатыва", 3, 6, 2},
		{"неучтивости", 4, 7, 2, true, "неучтивост", 4, 7, 2},
		{"устремлены", 2, 6, 1, true, "устремл", 2, 6, 1},
		{"голубых", 3, 5, 2, true, "голуб", 3, 5, 2},
		{"четвертый", 3, 6, 2, true, "четверт", 3, 6, 2},
		{"властному", 4, 8, 3, true, "властн", 4, 6, 3},
		{"стыдливо", 4, 7, 3, true, "стыдлив", 4, 7, 3},
		{"фигуре", 3, 5, 2, true, "фигур", 3, 5, 2},
		{"подкараулил", 3, 6, 2, true, "подкараул", 3, 6, 2},
		{"неосторожно", 4, 7, 2, true, "неосторожн", 4, 7, 2},
		{"особых", 2, 4, 1, true, "особ", 2, 4, 1},
		{"валерию", 3, 5, 2, true, "валери", 3, 5, 2},
		{"проходят", 4, 6, 3, true, "проход", 4, 6, 3},
		// False cases
		{"список", 4, 6, 3, false, "список", 4, 6, 3},
		{"погорячатся", 3, 5, 2, false, "погорячат", 3, 5, 2},
		{"пропеллер", 4, 6, 3, false, "пропеллер", 4, 6, 3},
		{"телефон", 3, 5, 2, false, "телефон", 3, 5, 2},
		{"кухарок", 3, 5, 2, false, "кухарок", 3, 5, 2},
		{"сошьет", 3, 6, 2, false, "сошьет", 3, 6, 2},
		{"проникся", 4, 6, 3, false, "проник", 4, 6, 3},
		{"стол", 4, 4, 3, false, "стол", 4, 4, 3},
		{"смахнет", 4, 7, 3, false, "смахнет", 4, 7, 3},
		{"торжеств", 3, 6, 2, false, "торжеств", 3, 6, 2},
		{"убор", 2, 4, 1, false, "убор", 2, 4, 1},
		{"тают", 4, 4, 2, false, "тают", 4, 4, 2},
		{"обнес", 2, 5, 1, false, "обнес", 2, 5, 1},
		{"промелькнет", 4, 6, 3, false, "промелькнет", 4, 6, 3},
		{"екатерингоф", 2, 4, 1, false, "екатерингоф", 2, 4, 1},
		{"упадет", 2, 4, 1, false, "упадет", 2, 4, 1},
		{"продавец", 4, 6, 3, false, "продавец", 4, 6, 3},
		{"рос", 3, 3, 2, false, "рос", 3, 3, 2},
		{"огляделся", 2, 5, 1, false, "оглядел", 2, 5, 1},
		{"пошепчут", 3, 5, 2, false, "пошепчут", 3, 5, 2},
		{"сошелся", 3, 5, 2, false, "сошел", 3, 5, 2},
		{"курс", 3, 4, 2, false, "курс", 3, 4, 2},
		{"вяз", 3, 3, 2, false, "вяз", 3, 3, 2},
		{"привлек", 4, 7, 3, false, "привлек", 4, 7, 3},
		{"копеек", 3, 6, 2, false, "копеек", 3, 6, 2},
		{"хочется", 3, 5, 2, false, "хочет", 3, 5, 2},
		{"лорнет", 3, 6, 2, false, "лорнет", 3, 6, 2},
		{"махнув", 3, 6, 2, false, "махнув", 3, 6, 2},
		{"дворянин", 4, 6, 3, false, "дворянин", 4, 6, 3},
		{"испанец", 2, 5, 1, false, "испанец", 2, 5, 1},
		{"выздоровел", 3, 6, 2, false, "выздоровел", 3, 6, 2},
		{"попыток", 3, 5, 2, false, "попыток", 3, 5, 2},
		{"пространств", 4, 8, 3, false, "пространств", 4, 8, 3},
		{"мус", 3, 3, 2, false, "мус", 3, 3, 2},
		{"сенат", 3, 5, 2, false, "сенат", 3, 5, 2},
		{"блин", 4, 4, 3, false, "блин", 4, 4, 3},
		{"жгут", 4, 4, 3, false, "жгут", 4, 4, 3},
		{"ликург", 3, 5, 2, false, "ликург", 3, 5, 2},
		{"люд", 3, 3, 2, false, "люд", 3, 3, 2},
		{"тьме", 4, 4, 4, false, "тьме", 4, 4, 4},
		{"сундучок", 3, 6, 2, false, "сундучок", 3, 6, 2},
		{"убьет", 2, 5, 1, false, "убьет", 2, 5, 1},
		{"чем", 3, 3, 2, false, "чем", 3, 3, 2},
		{"фонарик", 3, 5, 2, false, "фонарик", 3, 5, 2},
		{"холм", 3, 4, 2, false, "холм", 3, 4, 2},
		{"окинув", 2, 4, 1, false, "окинув", 2, 4, 1},
		{"границ", 4, 6, 3, false, "границ", 4, 6, 3},
		{"сам", 3, 3, 2, false, "сам", 3, 3, 2},
		{"овчинин", 2, 5, 1, false, "овчинин", 2, 5, 1},
		{"ков", 3, 3, 2, false, "ков", 3, 3, 2},
	}
	romance.RunStepTest(t, step1, testCases)

}

// Test step2
//
func Test_step2(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"незнани", 3, 6, 2, true, "незнан", 3, 6, 2},
		{"однообрази", 2, 6, 1, true, "однообраз", 2, 6, 1},
		{"развити", 3, 6, 2, true, "развит", 3, 6, 2},
		{"развити", 3, 6, 2, true, "развит", 3, 6, 2},
		{"милосерди", 3, 5, 2, true, "милосерд", 3, 5, 2},
		{"презрени", 4, 7, 3, true, "презрен", 4, 7, 3},
		{"заботливости", 3, 5, 2, true, "заботливост", 3, 5, 2},
		{"нмени", 4, 5, 3, true, "нмен", 4, 4, 3},
		{"соболезновани", 3, 5, 2, true, "соболезнован", 3, 5, 2},
		{"продолжени", 4, 6, 3, true, "продолжен", 4, 6, 3},
		{"религи", 3, 5, 2, true, "религ", 3, 5, 2},
		{"инструкци", 2, 7, 1, true, "инструкц", 2, 7, 1},
		{"выпуклости", 3, 5, 2, true, "выпуклост", 3, 5, 2},
		{"равнодуши", 3, 6, 2, true, "равнодуш", 3, 6, 2},
		{"равнодуши", 3, 6, 2, true, "равнодуш", 3, 6, 2},
		{"страдани", 5, 7, 4, true, "страдан", 5, 7, 4},
		{"страдани", 5, 7, 4, true, "страдан", 5, 7, 4},
		{"тягости", 3, 5, 2, true, "тягост", 3, 5, 2},
		{"необходимости", 4, 7, 2, true, "необходимост", 4, 7, 2},
		{"юри", 2, 3, 1, true, "юр", 2, 2, 1},
		{"фотографи", 3, 5, 2, true, "фотограф", 3, 5, 2},
		{"объяснени", 2, 5, 1, true, "объяснен", 2, 5, 1},
		{"объяснени", 2, 5, 1, true, "объяснен", 2, 5, 1},
		{"движени", 4, 6, 3, true, "движен", 4, 6, 3},
		{"движени", 4, 6, 3, true, "движен", 4, 6, 3},
		{"видени", 3, 5, 2, true, "виден", 3, 5, 2},
		{"подымани", 3, 5, 2, true, "подыман", 3, 5, 2},
		{"остервенени", 2, 5, 1, true, "остервенен", 2, 5, 1},
		{"организаци", 2, 5, 1, true, "организац", 2, 5, 1},
		{"давлени", 3, 6, 2, true, "давлен", 3, 6, 2},
		{"мономани", 3, 5, 2, true, "мономан", 3, 5, 2},
		{"заглави", 3, 6, 2, true, "заглав", 3, 6, 2},
		{"смягчени", 4, 7, 3, true, "смягчен", 4, 7, 3},
		{"морфи", 3, 5, 2, true, "морф", 3, 4, 2},
		{"оружи", 2, 4, 1, true, "оруж", 2, 4, 1},
		{"полупрезрени", 3, 5, 2, true, "полупрезрен", 3, 5, 2},
		{"офели", 2, 4, 1, true, "офел", 2, 4, 1},
		{"биени", 4, 5, 2, true, "биен", 4, 4, 2},
		{"биени", 4, 5, 2, true, "биен", 4, 4, 2},
		{"ази", 2, 3, 1, true, "аз", 2, 2, 1},
		{"приняти", 4, 6, 3, true, "принят", 4, 6, 3},
		{"отолщени", 2, 4, 1, true, "отолщен", 2, 4, 1},
		{"негодовани", 3, 5, 2, true, "негодован", 3, 5, 2},
		{"знамени", 4, 6, 3, true, "знамен", 4, 6, 3},
		{"благоговени", 4, 6, 3, true, "благоговен", 4, 6, 3},
		{"обожани", 2, 4, 1, true, "обожан", 2, 4, 1},
		{"евгени", 2, 5, 1, true, "евген", 2, 5, 1},
		{"комеди", 3, 5, 2, true, "комед", 3, 5, 2},
		{"показани", 3, 5, 2, true, "показан", 3, 5, 2},
		{"поручени", 3, 5, 2, true, "поручен", 3, 5, 2},
		// False cases
		{"устраня", 2, 6, 1, false, "устраня", 2, 6, 1},
		{"устраня", 2, 6, 1, false, "устраня", 2, 6, 1},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"ваш", 3, 3, 2, false, "ваш", 3, 3, 2},
		{"изум", 2, 4, 1, false, "изум", 2, 4, 1},
		{"изум", 2, 4, 1, false, "изум", 2, 4, 1},
		{"изум", 2, 4, 1, false, "изум", 2, 4, 1},
		{"изум", 2, 4, 1, false, "изум", 2, 4, 1},
		{"изум", 2, 4, 1, false, "изум", 2, 4, 1},
		{"постарет", 3, 6, 2, false, "постарет", 3, 6, 2},
		{"соснов", 3, 6, 2, false, "соснов", 3, 6, 2},
		{"развязыва", 3, 6, 2, false, "развязыва", 3, 6, 2},
		{"жуковск", 3, 5, 2, false, "жуковск", 3, 5, 2},
		{"разгоня", 3, 6, 2, false, "разгоня", 3, 6, 2},
		{"безнравственн", 3, 7, 2, false, "безнравственн", 3, 7, 2},
		{"безнравственн", 3, 7, 2, false, "безнравственн", 3, 7, 2},
		{"безнравственн", 3, 7, 2, false, "безнравственн", 3, 7, 2},
		{"безнравственн", 3, 7, 2, false, "безнравственн", 3, 7, 2},
		{"безнравственн", 3, 7, 2, false, "безнравственн", 3, 7, 2},
		{"неотеса", 4, 6, 2, false, "неотеса", 4, 6, 2},
		{"кудряв", 3, 6, 2, false, "кудряв", 3, 6, 2},
		{"протерт", 4, 6, 3, false, "протерт", 4, 6, 3},
		{"лавочк", 3, 5, 2, false, "лавочк", 3, 5, 2},
		{"лавочк", 3, 5, 2, false, "лавочк", 3, 5, 2},
		{"лавочк", 3, 5, 2, false, "лавочк", 3, 5, 2},
		{"лавочк", 3, 5, 2, false, "лавочк", 3, 5, 2},
		{"лавочк", 3, 5, 2, false, "лавочк", 3, 5, 2},
		{"лавочк", 3, 5, 2, false, "лавочк", 3, 5, 2},
		{"панорам", 3, 5, 2, false, "панорам", 3, 5, 2},
		{"панорам", 3, 5, 2, false, "панорам", 3, 5, 2},
		{"панорам", 3, 5, 2, false, "панорам", 3, 5, 2},
		{"торопливост", 3, 5, 2, false, "торопливост", 3, 5, 2},
		{"дщер", 4, 4, 3, false, "дщер", 4, 4, 3},
		{"заболева", 3, 5, 2, false, "заболева", 3, 5, 2},
		{"заболева", 3, 5, 2, false, "заболева", 3, 5, 2},
		{"заколдова", 3, 5, 2, false, "заколдова", 3, 5, 2},
		{"заколдова", 3, 5, 2, false, "заколдова", 3, 5, 2},
		{"додума", 3, 5, 2, false, "додума", 3, 5, 2},
		{"додума", 3, 5, 2, false, "додума", 3, 5, 2},
	}
	romance.RunStepTest(t, step2, testCases)

}

// Test step3
//
func Test_step3(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"подробност", 3, 6, 2, true, "подробн", 3, 6, 2},
		{"подробност", 3, 6, 2, true, "подробн", 3, 6, 2},
		{"подробност", 3, 6, 2, true, "подробн", 3, 6, 2},
		{"подробност", 3, 6, 2, true, "подробн", 3, 6, 2},
		{"подробност", 3, 6, 2, true, "подробн", 3, 6, 2},
		{"нерешительность", 3, 5, 2, true, "нерешительн", 3, 5, 2},
		{"принадлежност", 4, 6, 3, true, "принадлежн", 4, 6, 3},
		{"принадлежност", 4, 6, 3, true, "принадлежн", 4, 6, 3},
		{"восторженност", 3, 6, 2, true, "восторженн", 3, 6, 2},
		{"недоступност", 3, 5, 2, true, "недоступн", 3, 5, 2},
		{"болезненност", 3, 5, 2, true, "болезненн", 3, 5, 2},
		{"испакост", 2, 5, 1, true, "испак", 2, 5, 1},
		{"холодност", 3, 5, 2, true, "холодн", 3, 5, 2},
		{"холодност", 3, 5, 2, true, "холодн", 3, 5, 2},
		{"готовност", 3, 5, 2, true, "готовн", 3, 5, 2},
		{"ведомост", 3, 5, 2, true, "ведом", 3, 5, 2},
		{"ведомост", 3, 5, 2, true, "ведом", 3, 5, 2},
		{"непонятливост", 3, 5, 2, true, "непонятлив", 3, 5, 2},
		{"неровност", 3, 5, 2, true, "неровн", 3, 5, 2},
		{"обидчивост", 2, 4, 1, true, "обидчив", 2, 4, 1},
		{"сдержанность", 4, 7, 3, true, "сдержанн", 4, 7, 3},
		{"поспешность", 3, 6, 2, true, "поспешн", 3, 6, 2},
		{"справедливост", 5, 7, 4, true, "справедлив", 5, 7, 4},
		{"справедливост", 5, 7, 4, true, "справедлив", 5, 7, 4},
		{"неслиянност", 3, 7, 2, true, "неслиянн", 3, 7, 2},
		{"девственност", 3, 8, 2, true, "девственн", 3, 8, 2},
		{"насмешливост", 3, 6, 2, true, "насмешлив", 3, 6, 2},
		{"определенност", 2, 5, 1, true, "определенн", 2, 5, 1},
		{"безотчетност", 3, 5, 2, true, "безотчетн", 3, 5, 2},
		{"угодливост", 2, 4, 1, true, "угодлив", 2, 4, 1},
		{"брезгливост", 4, 8, 3, true, "брезглив", 4, 8, 3},
		{"непочтительност", 3, 5, 2, true, "непочтительн", 3, 5, 2},
		{"молодост", 3, 5, 2, true, "молод", 3, 5, 2},
		{"молодост", 3, 5, 2, true, "молод", 3, 5, 2},
		{"распущенност", 3, 6, 2, true, "распущенн", 3, 6, 2},
		{"внутренност", 4, 7, 3, true, "внутренн", 4, 7, 3},
		{"безоблачност", 3, 5, 2, true, "безоблачн", 3, 5, 2},
		{"ничтожност", 3, 6, 2, true, "ничтожн", 3, 6, 2},
		{"безрадостност", 3, 6, 2, true, "безрадостн", 3, 6, 2},
		{"неблагонадежност", 3, 6, 2, true, "неблагонадежн", 3, 6, 2},
		{"распорядительност", 3, 6, 2, true, "распорядительн", 3, 6, 2},
		{"распорядительност", 3, 6, 2, true, "распорядительн", 3, 6, 2},
		{"знаменитост", 4, 6, 3, true, "знаменит", 4, 6, 3},
		{"знаменитост", 4, 6, 3, true, "знаменит", 4, 6, 3},
		{"искренность", 2, 6, 1, true, "искренн", 2, 6, 1},
		{"заботливост", 3, 5, 2, true, "заботлив", 3, 5, 2},
		{"заботливост", 3, 5, 2, true, "заботлив", 3, 5, 2},
		{"заботливост", 3, 5, 2, true, "заботлив", 3, 5, 2},
		{"неловкост", 3, 5, 2, true, "неловк", 3, 5, 2},
		{"неловкост", 3, 5, 2, true, "неловк", 3, 5, 2},
		// False cases
		{"младенец", 4, 6, 3, false, "младенец", 4, 6, 3},
		{"перехитренн", 3, 5, 2, false, "перехитренн", 3, 5, 2},
		{"присватыва", 4, 7, 3, false, "присватыва", 4, 7, 3},
		{"непорядк", 3, 5, 2, false, "непорядк", 3, 5, 2},
		{"попят", 3, 5, 2, false, "попят", 3, 5, 2},
		{"меньшинств", 3, 7, 2, false, "меньшинств", 3, 7, 2},
		{"увар", 2, 4, 1, false, "увар", 2, 4, 1},
		{"покля", 3, 5, 2, false, "покля", 3, 5, 2},
		{"отвож", 2, 5, 1, false, "отвож", 2, 5, 1},
		{"туд", 3, 3, 2, false, "туд", 3, 3, 2},
		{"равнин", 3, 6, 2, false, "равнин", 3, 6, 2},
		{"ил", 2, 2, 1, false, "ил", 2, 2, 1},
		{"пройденн", 4, 7, 3, false, "пройденн", 4, 7, 3},
		{"будет", 3, 5, 2, false, "будет", 3, 5, 2},
		{"морг", 3, 4, 2, false, "морг", 3, 4, 2},
		{"выставля", 3, 6, 2, false, "выставля", 3, 6, 2},
		{"точк", 3, 4, 2, false, "точк", 3, 4, 2},
		{"мастерск", 3, 6, 2, false, "мастерск", 3, 6, 2},
		{"тих", 3, 3, 2, false, "тих", 3, 3, 2},
		{"пригорк", 4, 6, 3, false, "пригорк", 4, 6, 3},
		{"ожида", 2, 4, 1, false, "ожида", 2, 4, 1},
		{"широк", 3, 5, 2, false, "широк", 3, 5, 2},
		{"съеда", 4, 5, 3, false, "съеда", 4, 5, 3},
		{"проницательн", 4, 6, 3, false, "проницательн", 4, 6, 3},
		{"вычист", 3, 5, 2, false, "вычист", 3, 5, 2},
		{"погибнут", 3, 5, 2, false, "погибнут", 3, 5, 2},
		{"довод", 3, 5, 2, false, "довод", 3, 5, 2},
		{"прошумел", 4, 6, 3, false, "прошумел", 4, 6, 3},
		{"летевш", 3, 5, 2, false, "летевш", 3, 5, 2},
		{"подслужива", 3, 7, 2, false, "подслужива", 3, 7, 2},
		{"тургеневск", 3, 6, 2, false, "тургеневск", 3, 6, 2},
		{"выпрыгнул", 3, 6, 2, false, "выпрыгнул", 3, 6, 2},
		{"искусств", 2, 5, 1, false, "искусств", 2, 5, 1},
		{"подшиван", 3, 6, 2, false, "подшиван", 3, 6, 2},
		{"выгон", 3, 5, 2, false, "выгон", 3, 5, 2},
		{"вероломств", 3, 5, 2, false, "вероломств", 3, 5, 2},
		{"многодумн", 4, 6, 3, false, "многодумн", 4, 6, 3},
		{"примирен", 4, 6, 3, false, "примирен", 4, 6, 3},
		{"пескин", 3, 6, 2, false, "пескин", 3, 6, 2},
		{"рюмочк", 3, 5, 2, false, "рюмочк", 3, 5, 2},
		{"бархат", 3, 6, 2, false, "бархат", 3, 6, 2},
		{"блеснул", 4, 7, 3, false, "блеснул", 4, 7, 3},
		{"посторон", 3, 6, 2, false, "посторон", 3, 6, 2},
		{"единовременн", 2, 4, 1, false, "единовременн", 2, 4, 1},
		{"воплощен", 3, 6, 2, false, "воплощен", 3, 6, 2},
		{"сочиня", 3, 5, 2, false, "сочиня", 3, 5, 2},
		{"коробк", 3, 5, 2, false, "коробк", 3, 5, 2},
		{"съеш", 4, 4, 3, false, "съеш", 4, 4, 3},
		{"арестант", 2, 4, 1, false, "арестант", 2, 4, 1},
		{"шепнул", 3, 6, 2, false, "шепнул", 3, 6, 2},
	}
	romance.RunStepTest(t, step3, testCases)

}

// Test step4
//
func Test_step4(t *testing.T) {
	testCases := []romance.StepTestCase{
		{"раскольничь", 3, 6, 2, true, "раскольнич", 3, 6, 2},
		{"создань", 3, 6, 2, true, "создан", 3, 6, 2},
		{"доверенн", 3, 5, 2, true, "доверен", 3, 5, 2},
		{"уничтоженн", 2, 4, 1, true, "уничтожен", 2, 4, 1},
		{"опаленн", 2, 4, 1, true, "опален", 2, 4, 1},
		{"легкомысленн", 3, 6, 2, true, "легкомыслен", 3, 6, 2},
		{"поглощенн", 3, 6, 2, true, "поглощен", 3, 6, 2},
		{"дочерь", 3, 5, 2, true, "дочер", 3, 5, 2},
		{"каменн", 3, 5, 2, true, "камен", 3, 5, 2},
		{"извращенн", 2, 6, 1, true, "извращен", 2, 6, 1},
		{"печать", 3, 5, 2, true, "печат", 3, 5, 2},
		{"отчищенн", 2, 5, 1, true, "отчищен", 2, 5, 1},
		{"некупленн", 3, 5, 2, true, "некуплен", 3, 5, 2},
		{"отделенн", 2, 5, 1, true, "отделен", 2, 5, 1},
		{"покрашенн", 3, 6, 2, true, "покрашен", 3, 6, 2},
		{"наследственн", 3, 6, 2, true, "наследствен", 3, 6, 2},
		{"эмаль", 2, 4, 1, true, "эмал", 2, 4, 1},
		{"определенн", 2, 5, 1, true, "определен", 2, 5, 1},
		{"многопенн", 4, 6, 3, true, "многопен", 4, 6, 3},
		{"роднейш", 3, 6, 2, true, "родн", 3, 4, 2},
		{"чинн", 3, 4, 2, true, "чин", 3, 3, 2},
		{"смущенн", 4, 6, 3, true, "смущен", 4, 6, 3},
		{"единовременн", 2, 4, 1, true, "единовремен", 2, 4, 1},
		{"уставленн", 2, 5, 1, true, "уставлен", 2, 5, 1},
		{"наибесполезнейш", 4, 6, 2, true, "наибесполезн", 4, 6, 2},
		{"выкупленн", 3, 5, 2, true, "выкуплен", 3, 5, 2},
		{"софь", 3, 4, 2, true, "соф", 3, 3, 2},
		{"естественн", 2, 5, 1, true, "естествен", 2, 5, 1},
		{"замечательнейш", 3, 5, 2, true, "замечательн", 3, 5, 2},
		{"обеспокоенн", 2, 4, 1, true, "обеспокоен", 2, 4, 1},
		{"вражь", 4, 5, 3, true, "враж", 4, 4, 3},
		{"военн", 4, 5, 2, true, "воен", 4, 4, 2},
		{"поконченн", 3, 5, 2, true, "покончен", 3, 5, 2},
		{"помеченн", 3, 5, 2, true, "помечен", 3, 5, 2},
		{"отменн", 2, 5, 1, true, "отмен", 2, 5, 1},
		{"нетопленн", 3, 5, 2, true, "нетоплен", 3, 5, 2},
		{"веденн", 3, 5, 2, true, "веден", 3, 5, 2},
		{"несомненн", 3, 5, 2, true, "несомнен", 3, 5, 2},
		{"ревность", 3, 6, 2, true, "ревност", 3, 6, 2},
		{"освещенн", 2, 5, 1, true, "освещен", 2, 5, 1},
		{"добродушнейш", 3, 6, 2, true, "добродушн", 3, 6, 2},
		{"складенн", 5, 7, 4, true, "складен", 5, 7, 4},
		{"сиянь", 4, 5, 2, true, "сиян", 4, 4, 2},
		{"прославленн", 4, 7, 3, true, "прославлен", 4, 7, 3},
		{"поступь", 3, 6, 2, true, "поступ", 3, 6, 2},
		{"главнейш", 4, 7, 3, true, "главн", 4, 5, 3},
		{"птичь", 4, 5, 3, true, "птич", 4, 4, 3},
		{"коровь", 3, 5, 2, true, "коров", 3, 5, 2},
		{"ожерель", 2, 4, 1, true, "ожерел", 2, 4, 1},
		{"неестественн", 4, 7, 2, true, "неестествен", 4, 7, 2},
		// False cases
		{"украшен", 2, 5, 1, false, "украшен", 2, 5, 1},
		{"телеграфн", 3, 5, 2, false, "телеграфн", 3, 5, 2},
		{"унос", 2, 4, 1, false, "унос", 2, 4, 1},
		{"поездк", 4, 6, 2, false, "поездк", 4, 6, 2},
		{"перекрестк", 3, 5, 2, false, "перекрестк", 3, 5, 2},
		{"покомфортн", 3, 5, 2, false, "покомфортн", 3, 5, 2},
		{"своз", 4, 4, 3, false, "своз", 4, 4, 3},
		{"вытягив", 3, 5, 2, false, "вытягив", 3, 5, 2},
		{"казуистик", 3, 6, 2, false, "казуистик", 3, 6, 2},
		{"трог", 4, 4, 3, false, "трог", 4, 4, 3},
		{"предпочит", 4, 7, 3, false, "предпочит", 4, 7, 3},
		{"жгла", 4, 4, 4, false, "жгла", 4, 4, 4},
		{"возделыва", 3, 6, 2, false, "возделыва", 3, 6, 2},
		{"религ", 3, 5, 2, false, "религ", 3, 5, 2},
		{"бессердеч", 3, 6, 2, false, "бессердеч", 3, 6, 2},
		{"прохор", 4, 6, 3, false, "прохор", 4, 6, 3},
		{"сойдет", 3, 6, 2, false, "сойдет", 3, 6, 2},
		{"ораторствова", 2, 4, 1, false, "ораторствова", 2, 4, 1},
		{"прояв", 5, 5, 3, false, "прояв", 5, 5, 3},
		{"гербов", 3, 6, 2, false, "гербов", 3, 6, 2},
		{"понесл", 3, 5, 2, false, "понесл", 3, 5, 2},
		{"буфетчик", 3, 5, 2, false, "буфетчик", 3, 5, 2},
		{"поддержив", 3, 6, 2, false, "поддержив", 3, 6, 2},
		{"прячеш", 4, 6, 3, false, "прячеш", 4, 6, 3},
		{"загадочн", 3, 5, 2, false, "загадочн", 3, 5, 2},
		{"стряхнут", 5, 8, 4, false, "стряхнут", 5, 8, 4},
		{"пяточк", 3, 5, 2, false, "пяточк", 3, 5, 2},
		{"занов", 3, 5, 2, false, "занов", 3, 5, 2},
		{"солнечн", 3, 6, 2, false, "солнечн", 3, 6, 2},
		{"совестн", 3, 5, 2, false, "совестн", 3, 5, 2},
		{"совершеннолет", 3, 5, 2, false, "совершеннолет", 3, 5, 2},
		{"ищут", 2, 4, 1, false, "ищут", 2, 4, 1},
		{"сошедш", 3, 5, 2, false, "сошедш", 3, 5, 2},
		{"выпуст", 3, 5, 2, false, "выпуст", 3, 5, 2},
		{"молодец", 3, 5, 2, false, "молодец", 3, 5, 2},
		{"отрезв", 2, 5, 1, false, "отрезв", 2, 5, 1},
		{"топлив", 3, 6, 2, false, "топлив", 3, 6, 2},
		{"колодник", 3, 5, 2, false, "колодник", 3, 5, 2},
		{"присмирел", 4, 7, 3, false, "присмирел", 4, 7, 3},
		{"педагог", 3, 5, 2, false, "педагог", 3, 5, 2},
		{"распространя", 3, 7, 2, false, "распространя", 3, 7, 2},
		{"обиняк", 2, 4, 1, false, "обиняк", 2, 4, 1},
		{"дикар", 3, 5, 2, false, "дикар", 3, 5, 2},
		{"ят", 2, 2, 1, false, "ят", 2, 2, 1},
		{"слезьм", 4, 6, 3, false, "слезьм", 4, 6, 3},
		{"перешептыва", 3, 5, 2, false, "перешептыва", 3, 5, 2},
		{"раскаян", 3, 7, 2, false, "раскаян", 3, 7, 2},
		{"неисправим", 4, 8, 2, false, "неисправим", 4, 8, 2},
		{"цепля", 3, 5, 2, false, "цепля", 3, 5, 2},
		{"дворянин", 4, 6, 3, false, "дворянин", 4, 6, 3},
	}
	romance.RunStepTest(t, step4, testCases)
}

// Test a large set of words for which we know
// the correct stemmed form.
//
func Test_Vocabulary(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"странствуя", "странству"},
		{"маслом", "масл"},
		{"поражал", "поража"},
		{"тревожило", "тревож"},
		{"судырь", "судыр"},
		{"занять", "заня"},
		{"покраснеете", "покраснеет"},
		{"предприятием", "предприят"},
		{"несмелые", "несмел"},
		{"девицей", "девиц"},
		{"занятые", "занят"},
		{"гордостью", "гордост"},
		{"простояла", "простоя"},
		{"кофею", "коф"},
		{"пелагея", "пелаге"},
		{"нахальство", "нахальств"},
		{"содрогнется", "содрогнет"},
		{"провод", "провод"},
		{"оградясь", "оград"},
		{"пить", "пит"},
		{"юридическом", "юридическ"},
		{"нетерпелив", "нетерпел"},
		{"немом", "нем"},
		{"выгоняют", "выгоня"},
		{"рассчитаны", "рассчита"},
		{"спугнет", "спугнет"},
		{"оледенило", "оледен"},
		{"маловажной", "маловажн"},
		{"зарю", "зар"},
		{"боязливее", "боязлив"},
		{"солнечного", "солнечн"},
		{"подделывал", "подделыва"},
		{"откровенный", "откровен"},
		{"надворный", "надворн"},
		{"январе", "январ"},
		{"духоте", "духот"},
		{"хватили", "хват"},
		{"рысак", "рысак"},
		{"недосягаемым", "недосяга"},
		{"рассчитывая", "рассчитыв"},
		{"нетерпение", "нетерпен"},
		{"развившись", "разв"},
		{"подружились", "подруж"},
		{"подслуживается", "подслужива"},
		{"мочи", "моч"},
		{"подвел", "подвел"},
		{"расселись", "рассел"},
		{"ковша", "ковш"},
		{"сойду", "сойд"},
		{"дуни", "дун"},
		{"став", "став"},
		{"умен", "ум"},
		{"чистом", "чист"},
		{"утоли", "утол"},
		{"навзничь", "навзнич"},
		{"легкомысленно", "легкомыслен"},
		{"полнейшего", "полн"},
		{"простыню", "простын"},
		{"сообщили", "сообщ"},
		{"восьмом", "восьм"},
		{"плач", "плач"},
		{"оброка", "оброк"},
		{"полупросыпаясь", "полупросып"},
		{"ответил", "ответ"},
		{"плетнем", "плетн"},
		{"благодарна", "благодарн"},
		{"вялыми", "вял"},
		{"негодяй", "негодя"},
		{"предисловие", "предислов"},
		{"дворники", "дворник"},
		{"приезжали", "приезжа"},
		{"обвинить", "обвин"},
		{"батальоны", "батальон"},
		{"калека", "калек"},
		{"обольстительнее", "обольстительн"},
		{"коробочками", "коробочк"},
		{"избаловали", "избалова"},
		{"удалились", "удал"},
		{"архитектура", "архитектур"},
		{"барин", "барин"},
		{"конторе", "контор"},
		{"источнике", "источник"},
		{"сумасбродом", "сумасброд"},
		{"вырываются", "вырыва"},
		{"помнить", "помн"},
		{"ночуешь", "ночуеш"},
		{"мольбами", "мольб"},
		{"поставлен", "поставл"},
		{"привстала", "привста"},
		{"погрустив", "погруст"},
		{"бесчестным", "бесчестн"},
		{"погулять", "погуля"},
		{"материнское", "материнск"},
		{"клетку", "клетк"},
		{"стал", "стал"},
		{"келья", "кел"},
		{"харькоеве", "харькоев"},
		{"плечи", "плеч"},
		{"том", "том"},
		{"управления", "управлен"},
		{"хриплым", "хрипл"},
		{"привязанности", "привязан"},
		{"михайловским", "михайловск"},
		{"пересылались", "пересыла"},
		{"галлы", "галл"},
		{"именьишко", "именьишк"},
		{"обовьют", "обовьют"},
		{"упорная", "упорн"},
		{"плелся", "плел"},
		{"бешенстве", "бешенств"},
		{"непроницаема", "непроницаем"},
		{"туманных", "тума"},
		{"заблистала", "заблиста"},
		{"высшему", "высш"},
		{"спора", "спор"},
		{"подвернувшемуся", "подвернувш"},
		{"кротким", "кротк"},
		{"ногам", "ног"},
		{"синеющих", "синеющ"},
		{"батарею", "батар"},
		{"понятен", "понят"},
		{"фантазия", "фантаз"},
		{"ножами", "нож"},
		{"оборванной", "оборва"},
		{"неуклюжий", "неуклюж"},
		{"гадкий", "гадк"},
		{"кушать", "куша"},
		{"жирок", "жирок"},
		{"предавать", "предава"},
		{"вещим", "вещ"},
		{"плыл", "плыл"},
		{"добродетельного", "добродетельн"},
		{"сонных", "сон"},
		{"канал", "кана"},
		{"тучами", "туч"},
		{"деревню", "деревн"},
		{"оборванец", "оборванец"},
		{"хлада", "хлад"},
		{"толстый", "толст"},
		{"возни", "возн"},
		{"обезображеннее", "обезображен"},
		{"свертывают", "свертыва"},
		{"мутно", "мутн"},
		{"переговорено", "переговор"},
		{"начну", "начн"},
		{"суматохи", "суматох"},
		{"рыдающие", "рыда"},
		{"шлезвиг", "шлезвиг"},
		{"нерв", "нерв"},
		{"штопать", "штопа"},
		{"оттенок", "оттенок"},
		{"посолиднее", "посолидн"},
		{"щегольской", "щегольск"},
		{"поймет", "поймет"},
		{"ширится", "шир"},
		{"приискания", "приискан"},
		{"слабость", "слабост"},
		{"отпаривал", "отпарива"},
		{"многими", "мног"},
		{"артельщика", "артельщик"},
		{"детьми", "детьм"},
		{"понявший", "поня"},
		{"нянечка", "нянечк"},
		{"перелистывая", "перелистыв"},
		{"расплывается", "расплыва"},
		{"приглашает", "приглаша"},
		{"смотрель", "смотрел"},
		{"оспу", "осп"},
		{"удостоверившись", "удостовер"},
		{"кумиры", "кумир"},
		{"конфузливы", "конфузлив"},
		{"жилет", "жилет"},
		{"умеют", "умеют"},
		{"порфирий", "порфир"},
		{"решаете", "реша"},
		{"отмыли", "отм"},
		{"противоречило", "противореч"},
		{"ленивых", "ленив"},
		{"огородишком", "огородишк"},
		{"замерзнуть", "замерзнут"},
		{"приходится", "приход"},
		{"несчастный", "несчастн"},
		{"исправления", "исправлен"},
		{"высасывал", "высасыва"},
		{"подстилая", "подстил"},
		{"прекрасной", "прекрасн"},
		{"расстроенный", "расстроен"},
		{"подоконник", "подоконник"},
		{"образованный", "образова"},
		{"отжила", "отж"},
		{"поесть", "поест"},
		{"заслыша", "заслыш"},
		{"прибавилось", "прибав"},
		{"отправлении", "отправлен"},
		{"завитые", "завит"},
		{"недогадливый", "недогадлив"},
		{"изуродованный", "изуродова"},
		{"блаженство", "блаженств"},
		{"охмелел", "охмелел"},
		{"куртке", "куртк"},
	}
	for _, testCase := range testCases {
		result := Stem(testCase.in, true)
		if result != testCase.out {
			t.Errorf("Expected %v -> %v, but got %v", testCase.in, testCase.out, result)
		}
	}
}
