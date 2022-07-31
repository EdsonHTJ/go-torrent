

const torrentProgress = (t) => {
    return (
        <div className="Tinfo">
            <p>{t.name}</p>
            <p>---</p>
            <p>{t.downloaded * 100 / total}</p>
        </div>
    );
}