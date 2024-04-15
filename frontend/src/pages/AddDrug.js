import React, { useState } from 'react'
import DatePicker from "react-datepicker";
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

function AddDrug() {
    const [name, setName] = useState()
    const [quantity, setQuantity] = useState()
    const [location, setLocation] = useState()
    const [expiryDate, setExpiryDate] = useState()
    const navigate = useNavigate()

    const Submit = (e) => {
        e.preventDefault();
        axios.post("http://localhost:8080/drugs", {name, quantity, location, expiryDate})
            .then(result => { 
                console.log(result)
                navigate("/")
            })
            .catch(err => console.log(err))
    }

    return (
        <div className='d-flex vh-100 bg-primary justify-content-center align-items-center'>
            <div className='w-50 bg-white rounded p-3'>
                <form onSubmit={Submit}>
                    <h2>Add Drug</h2>
                    <div className='mb-2'>
                        <label htmlFor=''>Name</label>
                        <input type='text' placeholder='Enter Name' className='form-control'
                            onChange={(e) => setName(e.target.value)}
                        />
                    </div>
                    <div className='mb-2'>
                        <label htmlFor=''>Quantity</label>
                        <input type='text' placeholder='Enter Quantity' className='form-control'
                            onChange={(e) => setQuantity(e.target.value)}
                        />
                    </div>
                    <div className='mb-2'>
                        <label htmlFor=''>Location</label>
                        <input type='text' placeholder='Enter Location' className='form-control'
                            onChange={(e) => setLocation(e.target.value)}
                        />
                    </div>
                    <div className='mb-2'>
                        <label htmlFor=''>Expiry Date</label>
                        <input type='text' placeholder='Enter ' className='form-control'
                            onChange={(date) => setExpiryDate(date)}
                        />
                        {/* <DatePicker selected={expiryDate} onChange={(date) => setExpiryDate(date)} /> */}
                    </div>
                    <button className='btn btn-success'>Submit</button>
                </form>
            </div>
        </div>
    )
}

export default AddDrug