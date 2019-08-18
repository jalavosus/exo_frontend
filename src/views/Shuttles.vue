<template lang="pug">
  div.shuttles
    b-card(style="min-width:20rem; max-width:25rem;")
      b-card-body(title="Honey Badger's Back")
        b-form.form-parent(@submit="onSubmit" @reset="onReset" v-if="show")
          b-form-group#name(label="Please enter your name" label-for="name-input" label-align="left" label-class="boldlabel")
            b-form-input.input-field(
              id="name-input"
              v-model="form.name"
              type="text"
              required
              placeholder="Ploni Almoni")

          b-form-group#email(label="Please enter your email" label-for="email-input" label-align="left" label-class="boldlabel")
            b-form-input.input-field(
              id="email-input"
              v-model="form.email"
              type="email"
              required
              placeholder="ploni.almoni@gmail.com")

          b-form-group#direction(label="Select a direction" label-for="direction-dropdown" label-align="left" label-class="boldlabel")
            b-form-select.select-field(
              id="direction-dropdown"
              v-model="form.direction"
              :options="direction_list"
            )

          b-form-group#date(label="Select a date" label-for="datepicker" label-align="left" label-class="boldlabel")
            Datepicker#datepicker(v-model="form.date" :config="datepickerOptions" @dp-change="onDatepickerChange")

          b-form-group#timepicker(label="Select a time" label-for="time-select" label-align="left" label-class="boldlabel")
            div(v-if="form.direction === 'uptown'")
              b-form-select.select-field(
                v-if="checkWeekday(form.date)"
                id="time-select"
                v-model="form.time"
                :options="uptownWeekdayTimes")
              b-form-select.select-field(
                v-else-if="checkFriday(form.date)"
                id="time-select"
                v-model="form.time"
                :options="uptownFridayTimes")
              b-form-select.select-field(
                v-else-if="checkSaturday(form.date)"
                id="time-select"
                v-model="form.time")
              b-form-select.select-field(
                v-else
                id="time-select"
                v-model="form.time")
            div(v-else-if="form.direction === 'downtown'")
              b-form-select.select-field(
                v-if="checkWeekday(form.date)"
                id="time-select"
                v-model="form.time"
                :options="downtownWeekdayTimes")
              b-form-select.select-field(
                v-else-if="checkFriday(form.date)"
                id="time-select"
                v-model="form.time"
                :options="downtownFridayTimes")
              b-form-select.select-field(
                v-else
                id="time-select"
                v-model="form.time")
          b-button(type="submit" variant="primary" style="margin-bottom: -1.5em;" size="lg")
            | Book a shuttle!
</template>

<style scoped>
.shuttles {
  display: flex;
  justify-content: center;
}
.form-parent {
  margin-top: 2em;
}
.boldlabel {
  font-weight: 900;
}
</style>

<script>
  import Datepicker from 'vue-bootstrap-datetimepicker'
  import 'pc-bootstrap4-datetimepicker/build/css/bootstrap-datetimepicker.css'
  import { uptownWeekdayTimes, downtownWeekdayTimes } from "../shuttleTimes.js"
  import { uptownFridayTimes,  downtownFridayTimes }  from "../shuttleTimes.js"
  import Request from "request"

  export default {
    components: {
      Datepicker
    },
    data: function() {
      return {
        form: {
          name:       "",
          email:      "",
          direction:  "uptown",
          date:       new Date().toISOString().slice(0,10),
          time:       "12:00am"
        },
        show: true,
        direction_list: [
          { value: "uptown", text: "Uptown" }, { value: "downtown", text: "Downtown" }
        ],
        datepickerOptions: {
          format: "YYYY-MM-DD",
          useCurrent: true,
        },
        uptownWeekdayTimes:   uptownWeekdayTimes,
        downtownWeekdayTimes: downtownWeekdayTimes,
        uptownFridayTimes:    uptownFridayTimes,
        downtownFridayTimes:  downtownFridayTimes
      }
    },
    methods: {
      getDay(date) {
        var parsedDate = `${date}T00:00:00.000-0500`    // Purposefully using EST instead of EDT in all cases
        var obj = new Date(parsedDate)

        return obj.getDay()
      },
      checkWeekday(date) {
        var day = this.getDay(date)
        var isWeekday = true

        if (day === 0 || day === 5 || day === 6) {
          isWeekday = false
        }

        return isWeekday
      },
      checkFriday(date) {
        var day = this.getDay(date)

        return day === 5
      },
      checkSaturday(date) {
        var day = this.getDay(date)

        return day === 6
      },
      checkSunday(date) {
        var day = this.getDay(date)

        return day === 0
      },
      // eslint-disable-next-line
      onDatepickerChange(evt) {
        this.form.time = "12:00am"
      },
      onSubmit(evt) {
        evt.preventDefault()
        let formData = JSON.parse(JSON.stringify(this.form))
        Request.post(
          "http://localhost:8080/shuttles",
          { form: formData },
          // eslint-disable-next-line
          function (error, response, body) {
            if (response.statusCode == 200) {
              // eslint-disable-next-line
              console.log("Hell yeah")
            }
          }
        )
      },
      onReset(evt) {
        evt.preventDefault()
        // Reset our form values
        this.form.name = ""
        this.form.email = ""
        this.form.direction = "uptown"
        this.form.date = new Date().toISOString().slice(0, 10)
        this.form.checked = []
        // Trick to reset/clear native browser form validation state
        this.show = false
        this.$nextTick(() => {
          this.show = true
        })
      }
    }
  }
</script>
