import { Component, Vue, Prop, Emit } from "vue-property-decorator";

@Component({})
export default class DatePickerComponent extends Vue {

  @Prop({ type: String })
  label!: string;

  selectedDate= "";
  showCalendar= false;

  @Emit()
  selected(dateStr: string) {
    console.log(`emitted ${dateStr} from DatePickerComponent`);
  }

  onEndSelectedDate() {
    this.showCalendar= false;
    this.selected(this.selectedDate);
  }

  constructor() {
    super();
    //this.labelText = "";
  }

  // created() {
  // }

  // mounted() {
  // }
}
