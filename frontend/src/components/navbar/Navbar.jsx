import React from 'react';
import { BrowserRouter, Link } from 'react-router-dom';
import './Navbar.css';
import logo from '../../assets/img/logo.svg'

const Navbar = () => {
  return (
    <>
      <div className="navbar">
        <div className="container">
          <BrowserRouter>
            {/* 
              Ссылка с роутером сделана таким образом только для того, чтобы небыло перезагрузки при клике.
              Возможно, это избыточно, и ссылку можно было не делать. Ведь сайт состоит из единственной главной страницы.
              Пусть будет
            */}
            <Link to="/"><img className="navbar__logo" alt="Конвертирование видео" src={logo} /></Link>
          </BrowserRouter>
        </div>
      </div>
    </>
  );
}

export default Navbar;
