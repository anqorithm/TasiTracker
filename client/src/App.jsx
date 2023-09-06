import React, { useState, useEffect } from 'react';
import CompanyCard from './components/CompanyCard';
import Header from './components/Header';
import Footer from './components/Footer';
function App() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('http://127.0.0.1:8080/api/v1/companies?limit=1000')
      .then(response => response.json())
      .then(data => {
        setData(data.data);
        setLoading(false);
      })
      .catch(error => {
        console.error("Error fetching data:", error);
        setLoading(false);
      });
  }, []);

  if (loading) return <div>Loading...</div>;

  return (
    <div>
      <Header />
      <div className="container mx-auto p-4" dir="rtl">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {data.map(company => (
            <CompanyCard key={company.guid} company={company} />
          ))}
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default App;
