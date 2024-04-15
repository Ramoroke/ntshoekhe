import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'

function Home() {
    const [drugs, setDrugs] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:8080/drugs")
            .then( result => {
                setDrugs(result.data.items)
                console.log(drugs)
            })
            .catch(err => console.log(err))
    })

    const handleDelete = (id) => {
        axios.delete('http://localhost:8080/drugs'+id)
            .then(res => {
                console.log(res)
                window.location.reload()
            })
            .catch(err => console.log(err))
    }

  return (
    <div>
        <div className='d-flex vh-100 bg-primary justify-content-center align-items-center'>
            <div className='w-50 bg-white rounded p-3'>
                <Link to="/create" className='btn btn-success' >Add Drug + </Link>
                <table className='table'>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Quantity</th>
                            <th>Location</th>
                            <th>ExpiryDate</th>
                        </tr>
                    </thead>
                    <tbody>
                        {drugs.map((drug) => {
                            return (
                                <tr>
                                <td>{drug.name}</td>
                                <td>{drug.quantity}</td>
                                <td>{drug.location}</td>
                                <td>{drug.expiry_date}</td>
                                <td>
                                    <Link to='/update' className='btn btn-success'>Edit</Link>
                                    <button className='btn btn-danger'
                                        onClick={(e) => handleDelete(drug.id)}
                                    >Delete</button>
                                </td>
                                </tr>
                            )
                        })}
                    </tbody>
                </table> 
            </div>
        </div>
    </div>
  )
}

export default Home
