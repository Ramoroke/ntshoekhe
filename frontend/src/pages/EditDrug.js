import React, {useState, useEffect} from 'react'
import {useParams, useNavigate} from 'react-router-dom'
import axios from 'axios';

function EditDrug() {
    const {id} = useParams();
    const [name, setName] = useState()
    const [quantity, setQuantity] = useState()
    const [location, setLocation] = useState()
    const [expiryDate, setExpiryDate] = useState()
    const navigate = useNavigate()

    useEffect(() => {
        axios.get("http://localhost:8080/drugs/"+ id)
            .then(result => {
                setName(result.data.items.name)
                setQuantity(result.data.items.quantity)
                setLocation(result.data.items.location)
                setExpiryDate(result.data.items.expiry_date)
            })
            .catch(err => console.log(err))
    })

    const Update = (e) => {
        e.preventDefault();
        axios.put("http://localhost:8080/drugs"+id, {name, quantity, location, expiryDate})
            .then(result => { 
                console.log(result)
                navigate("/")
            })
            .catch(err => console.log(err))
    }

    return (
        <div className='d-flex vh-100 bg-primary justify-content-center align-items-center'>
            <div className='w-50 bg-white rounded p-3'>
                <form onSubmit={Update}>
                    <h2>Edit Drug</h2>
                    <div className='mb-2'>
                        <label htmlFor=''>Name</label>
                        <input type='text' placeholder='Enter Name' className='form-control'
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                    </div>
                    <div className='mb-2'>
                        <label htmlFor=''>Quantity</label>
                        <input type='text' placeholder='Enter Quantity' className='form-control'
                         value={quantity}
                         onChange={(e) => setQuantity(e.target.value)}
                        />
                    </div>
                    <div className='mb-2'>
                        <label htmlFor=''>Location</label>
                        <input type='text' placeholder='Enter Location' className='form-control'
                            value={location}
                            onChange={(e) => setLocation(e.target.value)}
                        />
                    </div>
                    <div className='mb-2'>
                        <label htmlFor=''>Expiry Date</label>
                        <input type='text' placeholder='Enter ' className='form-control'
                            value={expiryDate}
                            onChange={(date) => setExpiryDate(date)}
                        />
                    </div>
                    <button className='btn btn-success'>Update</button>
                </form>
            </div>
        </div>
    )
}

export default EditDrug