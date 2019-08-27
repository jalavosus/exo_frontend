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
              @change="onDirectionPickerChange"
            )

          b-form-group#date(label="Select a date" label-for="datepicker" label-align="left" label-class="boldlabel")
            Datepicker#datepicker(v-model="form.date" :config="datepickerOptions" @dp-change="onDatepickerChange")

          b-form-group#timepicker(label="Select a time" label-for="time-select" label-align="left" label-class="boldlabel")
            b-form-select.select-field(
              id="time-select"
              v-model="form.time"
              :options="shuttleTimes")

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
  /* eslint-disable */
  import Datepicker from 'vue-bootstrap-datetimepicker'
  import 'pc-bootstrap4-datetimepicker/build/css/bootstrap-datetimepicker.css'
  import { uptownWeekdayTimes, downtownWeekdayTimes } from "../shuttleTimes.js"
  import { uptownFridayTimes,  downtownFridayTimes }  from "../shuttleTimes.js"
  import Request from "request-promise"
  import Cache from "node-cache"

  let CacheMan = new Cache()

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
          time:       "12:00 AM"
        },
        show: true,
        direction_list: [
          { value: "uptown", text: "Uptown" }, { value: "downtown", text: "Downtown" }
        ],
        datepickerOptions: {
          format: "YYYY-MM-DD",
          useCurrent: true,
        },
        shuttleTimes: [],
      }
    },
    created() {
      console.log("Test")
      this.getShuttleTimes(this.form.date, this.form.direction)
    },
    methods: {
      getDay(date) {
        var parsedDate = `${date}T00:00:00.000-0500`    // Purposefully using EST instead of EDT in all cases
        var obj = new Date(parsedDate)

        return obj.getDay()
      },
      getShuttleTimes(date, direction) {
        var direc = 0
        switch (direction) {
          case "uptown":    direc = 1; break
          case "downtown":  direc = 2; break
        }

        let cacheStr = `times-${date}-${direction}`
        let cacheCheck = CacheMan.get(cacheStr)
        if (cacheCheck != null) {
          this.shuttleTimes = cacheCheck
          return
        }

        let self = this

        Request.get(`https://backend.exo.fish/shuttleTimes?direction=${direc}&date=${date}`)
          .then(function(result) {
            let newTimes = JSON.parse(result)["available_times"]
            self.shuttleTimes = newTimes
            CacheMan.set(cacheStr, newTimes)
          })
          .catch(function(error) {
            console.log(error)
          })
      },
      onDatepickerChange(evt) {
        this.form.time = "12:00 AM"
        this.getShuttleTimes(this.form.date, this.form.direction)
      },
      onDirectionPickerChange(newDirection) {
        this.getShuttleTimes(this.form.date, newDirection)
      },
      onSubmit(evt) {
        evt.preventDefault()
        let formData = JSON.parse(JSON.stringify(this.form))
        Request.post(
          "https://backend.exo.fish/shuttles",
          { form: formData },
          // eslint-disable-next-line
          function (error, response, body) {
            if (response.statusCode == 200) {
              // console.log("Hell yeah")
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
