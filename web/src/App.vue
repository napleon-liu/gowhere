<template>
<div id="app">
    <el-tabs v-model="activeTab">
        <el-tab-pane label="管理员" name="manage">
            <!-- 这里是管理航班、大巴车、宾馆房间和客户数据的代码 -->
            <el-tabs v-model="activeTab">
                <el-tab-pane label="航班" name="flight">
                    <el-table :data="flights" style="width: 30%">
                        <el-table-column prop="flightNum" label="航班号"></el-table-column>
                        <el-table-column prop="fromCity" label="出发地"></el-table-column>
                        <el-table-column prop="arivCity" label="目的地"></el-table-column>
                        <el-table-column prop="numSeats" label="座位数"></el-table-column>
                        <el-table-column prop="numAvail" label="剩余座位数"></el-table-column>
                        <el-table-column prop="price" label="价格"></el-table-column>
                        <el-table-column label="操作">
                            <el-button type="primary" size="small">编辑</el-button>
                            <el-button type="danger" size="small">删除</el-button>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="大巴车" name="bus">
                    <el-table :data="buses" style="width: 100%">
                        <el-table-column prop="location" label="城市"></el-table-column>
                        <el-table-column prop="numSeats" label="座位数"></el-table-column>
                        <el-table-column prop="numAvail" label="剩余座位数"></el-table-column>
                        <el-table-column prop="price" label="价格"></el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="宾馆房间" name="hotel">
                    <el-table :data="hotels" style="width: 100%">
                        <el-table-column prop="location" label="城市"></el-table-column>
                        <el-table-column prop="numRooms" label="房间总数"></el-table-column>
                        <el-table-column prop="numAvail" label="可用房间数"></el-table-column>
                        <el-table-column prop="price" label="价格"></el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="客户" name="customer">
                    <el-table :data="customers" style="width: 100%">
                        <el-table-column prop="custID" label="ID"></el-table-column>
                        <el-table-column prop="custName" label="姓名"></el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="预定资源" name="book">
                    <el-form ref="form" :model="form" label-width="120px">
                        <el-form-item label="航班号">
                            <el-input></el-input>
                        </el-form-item>
                        <el-form-item label="大巴车号">
                            <el-input></el-input>
                        </el-form-item>
                        <el-form-item label="宾馆房间号">
                            <el-input></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click="submitForm('form')">提交</el-button>
                        </el-form-item>
                    </el-form>
                </el-tab-pane>
            </el-tabs>
        </el-tab-pane>
        <el-tab-pane label="用户" name="book">
            <!-- 这里是管理航班、大巴车、宾馆房间和客户数据的代码 -->
            <el-tabs v-model="activeTab">
                <el-tab-pane label="航班" name="flight">
                    <el-table :data="flights" style="width: 30%">
                        <el-table-column prop="flightNum" label="航班号"></el-table-column>
                        <el-table-column prop="fromCity" label="出发地"></el-table-column>
                        <el-table-column prop="arivCity" label="目的地"></el-table-column>
                        <el-table-column prop="numSeats" label="座位数"></el-table-column>
                        <el-table-column prop="numAvail" label="剩余座位数"></el-table-column>
                        <el-table-column prop="price" label="价格"></el-table-column>
                        <el-table-column label="操作">
                            <el-button type="primary" size="small">预定</el-button>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="大巴车" name="bus">
                    <el-table :data="buses" style="width: 30%">
                        <el-table-column prop="location" label="城市"></el-table-column>
                        <el-table-column prop="numSeats" label="座位数"></el-table-column>
                        <el-table-column prop="numAvail" label="剩余座位数"></el-table-column>
                        <el-table-column prop="price" label="价格"></el-table-column>
                        <el-table-column label="操作">
                            <el-button type="primary" size="small">预定</el-button>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="宾馆房间" name="hotel">
                    <el-table :data="hotels" style="width: 30%">
                        <el-table-column prop="location" label="城市"></el-table-column>
                        <el-table-column prop="numRooms" label="房间总数"></el-table-column>
                        <el-table-column prop="numAvail" label="可用房间数"></el-table-column>
                        <el-table-column prop="price" label="价格"></el-table-column>
                        <el-table-column label="操作">
                            <el-button type="primary" size="small">预定</el-button>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
            </el-tabs>
        </el-tab-pane>
    </el-tabs>
</div>
</template>

<script>
// import { ElTable, ElTableColumn } from 'element-plus';
export default {
    name: 'App',
    components: {
        // ElTable,
        // ElTableColumn
    },
    data() {
        return {
            flights: [],
            buses: [],
            hotels: [],
            customers: []
        }
    },
    created() {
        // Fetch flight data from API and assign it to this.flights
        fetch('http://localhost:80/travel_booking/v1/flight')
            .then(response => response.json())
            .then(data => {
                this.flights = data.Result.flights;
            })
            .catch(error => {
                console.error('Error fetching flight data:', error);
            });
        // Fetch bus data from API and assign it to this.buses api: http://localhost:80/travel_booking/v1/bus
        fetch('http://localhost:80/travel_booking/v1/bus')
            .then(response => response.json())
            .then(data => {
                this.buses = data.Result.buses;
            })
            .catch(error => {
                console.error('Error fetching bus data:', error);
            });

        // Fetch hotel data from API and assign it to this.hotels api: http://localhost:80/travel_booking/v1/hotel
        fetch('http://localhost:80/travel_booking/v1/hotel')
            .then(response => response.json())
            .then(data => {
                this.hotels = data.Result.hotels;
            })
            .catch(error => {
                console.error('Error fetching hotel data:', error);
            });

        // Fetch customer data from API and assign it to this.customers api: http://localhost:80/travel_booking/v1/customer
        fetch('http://localhost:80/travel_booking/v1/customer')
            .then(response => response.json())
            .then(data => {
                this.customers = data.Result.customers;
            })
            .catch(error => {
                console.error('Error fetching customer data:', error);
            });
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    alert('submit!');
                    // 在这里添加你的预定代码
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        }
    }
}
</script>
