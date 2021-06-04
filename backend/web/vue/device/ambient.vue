<template>
	<div>
		<!-- ============================================= -->
		<b-row>
			<b-col
				cols="12"
				class="color-index border add-button-style"
			>
				<icon-device-aircon/>
			</b-col>
		</b-row>

		<b-row>
			<b-col
				cols="12"
				:class="[commonClass, methodAirconClass('airconStop')]"
				@click="methodSendCommand($event)"
				data-object="aircon"
				data-value="airconStop"
			>
				<icon-power
					data-object="aircon"
					data-value="airconStop"
				/>
			</b-col>
		</b-row>

		<b-row>
			<b-col
				cols="4"
				:class="[commonClass, methodAirconClass('airconWarm')]"
				@click="methodSendCommand($event)"
				data-object="aircon"
				data-value="airconWarm"
			>
				<icon-warm
					data-object="aircon"
					data-value="airconWarm"
				/>
			</b-col>

			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-object="aircon"
				data-key="temperature"
				data-value="1"
				@click="methodChangeTemperature($event)"
			>
				<icon-up
					data-object="aircon"
					data-key="temperature"
					data-value="1"
				/>
			</b-col>

			<b-col
				cols="4"
				:class="[commonClass, methodAirconClass('airconCool')]"
				@click="methodSendCommand($event)"
				data-object="aircon"
				data-value="airconCool"
			>
				<icon-cool
					data-object="aircon"
					data-value="airconCool"
				/>
			</b-col>
		</b-row>

		<b-row>
			<b-col
				cols="4"
				:class="[commonClass, methodAirconClass('airconWarm')]"
				@click="methodSendCommand($event)"
				data-object="aircon"
				data-value="airconWarm"
			>
			<span
				data-object="aircon"
				data-value="airconWarm"
			>
				{{airconWarm}}
			</span>
			</b-col>

			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-object="aircon"
				data-key="temperature"
				data-value="-1"
				@click="methodChangeTemperature($event)"

			>
				<icon-down
					data-object="aircon"
					data-key="temperature"
					data-value="-1"
				/>
			</b-col>

			<b-col
				cols="4"
				:class="[commonClass, methodAirconClass('airconCool')]"
				@click="methodSendCommand($event)"
				data-object="aircon"
				data-value="airconCool"
			>
			<span
				data-object="aircon"
				data-value="airconCool"
			>
				{{airconCool}}
			</span>
			</b-col>
		</b-row>

		<br>

		<!-- ============================================= -->
		<b-row>
			<b-col
				cols="12"
				class="color-index border add-button-style"
			>
				<icon-device-light/>
			</b-col>
		</b-row>

		<b-row>
			<b-col
				cols="4"
				:class="[commonClass, methodLightClass('lightPower')]"
				data-object="light"
				data-value="lightPower"
				@click="methodSendCommand($event)"
			>
				<icon-power
					data-object="light"
					data-value="lightPower"
				/>
			</b-col>

			<b-col
				cols="4"
				:class="[commonClass, methodLightClass('lightNight')]"
				data-object="light"
				data-value="lightNight"
				@click="methodSendCommand($event)"
			>
				<icon-night
					data-object="light"
					data-value="lightNight"
				/>
			</b-col>

			<b-col
				cols="4"
				:class="[commonClass, methodLightClass('lightFull')]"
				data-object="light"
				data-value="lightFull"
				@click="methodSendCommand($event)"
			>
				<icon-full
					data-object="light"
					data-value="lightFull"
				/>
			</b-col>
		</b-row>
	</div>
</template>


<script>
module.exports = {
	components: {
		'icon-device-aircon': httpVueLoader('web/vue/icon/device-aircon.vue'),
		'icon-device-light': httpVueLoader('web/vue/icon/device-light.vue'),

		'icon-cool': httpVueLoader('web/vue/icon/temperature-cool.vue'),
		'icon-warm': httpVueLoader('web/vue/icon/temperature-warm.vue'),
		'icon-up': httpVueLoader('web/vue/icon/chevron-up.vue'),
		'icon-down': httpVueLoader('web/vue/icon/chevron-down.vue'),

		'icon-power': httpVueLoader('web/vue/icon/common-powerOnOff.vue'),
		'icon-night': httpVueLoader('web/vue/icon/light-night.vue'),
		'icon-full': httpVueLoader('web/vue/icon/light-full.vue'),

	},
	methods: {
		methodAirconClass (value) {
			return value == this.airconStatus ? 'btn-primary' : 'btn-secondary';
		},

		methodLightClass (value) {
			return value == this.lightStatus ? 'btn-primary' : 'btn-secondary';
		},

		methodSendCommand: function (e) {
			let args = {};
			args['Object'] = e.target.dataset.object;
			args['Key'] = e.target.dataset.key == undefined ? this.defaultKey : e.target.dataset.key;
			args['Value'] = e.target.dataset.value;
			this.$emit('send-command', args);
		},

		methodChangeTemperature: function(e) {
			if ( this.airconStatus == 'airconStop' ){ return }

			let newTemp = 0;
			let currentValue = e.target.dataset.value - 0;
			let key = "";
			if ( this.airconStatus == 'airconWarm' ){
				newTemp = this.airconWarm + currentValue;
				key = "warmTemperature"
			}

			if ( this.airconStatus == 'airconCool' ){
				newTemp = this.airconCool + currentValue;
				key = "coolTemperature"
			}

			if (
				this.airconStatus == 'airconWarm' &&
				(newTemp < this.warmLimitLower || this.warmLimitHigher < newTemp )
			){ return }
			if (
				this.airconStatus == 'airconCool' &&
				(newTemp < this.coolLimitLower || this.coolLimitHigher < newTemp )
			){ return }

			let args = {};
			args['Object'] = e.target.dataset.object;
			args['Key'] = key;
			args['Value'] = String(newTemp);
			this.$emit('send-command', args);
		},
	},

	props: {
		'commonClass': {type: String, default: ''},
		'lightStatus': {type: String, default: ''},
		'airconStatus': {type: String, default: ''},
		'airconWarm': {type: Number , default: 0},
		'airconCool': {type: Number , default: 0},

},

	data: function() {
		return {
			warmLimitHigher: 30,
			warmLimitLower: 20,
			coolLimitHigher: 25,
			coolLimitLower: 15,
			defaultKey :"status",
		}
	}
}




